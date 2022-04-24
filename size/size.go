// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package size provides type Size to keep, marshal and unmarshal times-byte size values.
package size

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/bits"
	"reflect"
	"strconv"

	"go.lstv.dev/util/constraint"
	"go.lstv.dev/util/internal"
)

var (
	// DisableMarshalTextUnit allows disabling unit specification at Size.MarshalText.
	DisableMarshalTextUnit = false

	// DisableMarshalJSONStringForm allows disabling string form at Size.MarshalJSON.
	DisableMarshalJSONStringForm = false

	// DisableMarshalJSONObjectForm allows disabling object form at Size.MarshalJSON.
	DisableMarshalJSONObjectForm = false
)

// Size represents size in bytes.
// If unit is not present, number is always represented as value in bytes.
type Size uint64

// New creates new Size value with specified unit.
func New[N constraint.Numbers](value N, unit string) (Size, error) {
	s, err := newSize(value, unit)
	if err != nil {
		return 0, fmt.Errorf("size.New: %w", err)
	}
	return s, nil
}

// Shorten returns the biggest unit as is possible for value without rounding.
// Returned unit is always valid and binary (1024^x).
// Example: For Size(1024) is returned (1, "KiB"), but for Size(1025) is returned (1025, "B").
func (s Size) Shorten() (value uint64, unit string) {
	if s == 0 {
		return 0, Byte
	}
	v := uint64(s)
	for _, u := range shortenUnits {
		// try to divide by 1024 without remainder
		if (v & 0x3ff) != 0 {
			return v, u
		}
		v >>= 10
	}
	// maximum unit for uint64 is Exbibyte
	return v, Exbibyte
}

// BytesJSONNumber returns size as json.Number in bytes.
func (s Size) BytesJSONNumber() json.Number {
	return json.Number(strconv.FormatUint(uint64(s), 10))
}

// BytesString returns size as string in bytes.
func (s Size) BytesString() string {
	return strconv.FormatUint(uint64(s), 10)
}

// MarshalText converts size to text.
// If DisableMarshalTextUnit is false, Formatter is used.
// Otherwise, text form in bytes without unit is returned.
//
// See also DisableMarshalTextUnit.
func (s Size) MarshalText() ([]byte, error) {
	b, err := s.marshalText()
	if err != nil {
		return nil, fmt.Errorf("size.Size.MarshalText: %w", err)
	}
	return b, nil
}

// UnmarshalText using global Parser function.
// DefaultRule affects UnmarshalText behavior.
func (s *Size) UnmarshalText(data []byte) error {
	v, err := Parser(data, DefaultRule&ruleUnmarshalTextMask)
	if err != nil {
		return fmt.Errorf("size.Size.UnmarshalText: %w", err)
	}
	*s = v
	return nil
}

// MarshalJSON converts size to JSON value.
// If DisableMarshalJSONObjectForm is false, JSON object form is used.
// Otherwise, if DisableMarshalJSONStringForm is false, Size.MarshalText is used as string form.
// Otherwise, JSON number form is used (in bytes).
//
// Example of JSON object form for Size(1024):
//   {
//     "value": 1,
//     "unit": "KiB"
//   }
//
// See also DisableMarshalJSONObjectForm and DisableMarshalJSONStringForm.
func (s Size) MarshalJSON() ([]byte, error) {
	if !DisableMarshalJSONObjectForm {
		return s.marshalJSONObject(), nil
	}

	if !DisableMarshalJSONStringForm {
		b, err := s.marshalText()
		if err != nil {
			return nil, fmt.Errorf("size.Size.MarshalJSON: %w", err)
		}
		l := len(b)
		b = append(b, ` "`...)
		copy(b[1:l+1], b[:l])
		b[0] = '"'
		return b, nil
	}

	return strconv.AppendUint(nil, uint64(s), 10), nil
}

// UnmarshalJSON using global Parser function.
// DefaultRule affects UnmarshalJSON behavior.
func (s *Size) UnmarshalJSON(data []byte) error {
	v, err := Parser(data, DefaultRule)
	if err != nil {
		return fmt.Errorf("size.Size.UnmarshalJSON: %w", err)
	}
	*s = v
	return nil
}

// PrettyHTML formats size for HTML template.
// If Formatter returns error, PrettyHTML panics.
func (s Size) PrettyHTML() template.HTML {
	b := []byte(nil)
	b, err := Formatter(b, s, FormatPretty|FormatHTML)
	if err != nil {
		panic(err)
	}
	return template.HTML(b)
}

// PrettyString formats size for user output.
// If Formatter returns error, PrettyString panics.
func (s Size) PrettyString() string {
	b := []byte(nil)
	b, err := Formatter(b, s, FormatPretty)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// String formats size for string output.
// If Formatter returns error, String returns same value as BytesString.
func (s Size) String() string {
	b := []byte(nil)
	b, err := Formatter(b, s, 0)
	if err != nil {
		return strconv.FormatUint(uint64(s), 10)
	}
	return string(b)
}

// Bytes returns size value in bytes as specified data type.
// Returned ok is true if size value in bytes is suitable for specified data type.
func Bytes[N constraint.Numbers](s Size) (value N, ok bool) {
	switch k := internal.Kind(N(0)); k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uint64(s) <= uint64(internal.Max[N](k)) {
			return N(s), true
		}
	case reflect.Float32, reflect.Float64:
		if uint64(s) == uint64(N(s)) {
			return N(s), true
		}
	}
	return 0, false
}

func newSize[N constraint.Numbers](value N, unit string) (Size, error) {
	if value == 0 {
		if _, ok := zeroUnits[unit]; !ok {
			return 0, newInvalidUnitError(unit)
		}
		return 0, nil
	}
	if value < 0 || N(uint64(value)) != value {
		return 0, newInvalidValueError(value, unit)
	}
	if unit == "" {
		return Size(value), nil
	}
	n, ok := unitToValues[unit]
	if !ok {
		return 0, newInvalidUnitError(unit)
	}
	hi, lo := bits.Mul64(uint64(value), n)
	if hi != 0 {
		return 0, newInvalidValueError(value, unit)
	}
	return Size(lo), nil
}

func (s Size) marshalText() ([]byte, error) {
	if DisableMarshalTextUnit {
		return strconv.AppendUint(nil, uint64(s), 10), nil
	}

	b := []byte(nil)
	b, err := Formatter(b, s, 0)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s Size) marshalJSONObject() []byte {
	value, unit := s.Shorten()
	b := make([]byte, 0, 32)
	b = append(b, `{"`+ObjectKeyValue+`":`...)
	b = strconv.AppendUint(b, value, 10)
	b = append(b, `,"`+ObjectKeyUnit+`":"`...)
	b = append(b, unit...)
	b = append(b, `"}`...)
	return b
}
