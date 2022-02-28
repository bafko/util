// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package size provides type Size to keep, marshal and unmarshal times-byte size values.
package size

import (
	"encoding/json"
	"html/template"
	"math"
	"math/bits"
	"strconv"
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
func New(value uint64, unit string) (Size, error) {
	if value == 0 {
		if _, ok := zeroUnits[unit]; !ok {
			return 0, newInvalidUnitError(unit)
		}
		return 0, nil
	}
	if unit == "" {
		return Size(value), nil
	}
	n, ok := unitToValues[unit]
	if !ok {
		return 0, newInvalidUnitError(unit)
	}
	hi, lo := bits.Mul64(value, n)
	if hi != 0 {
		return 0, newInvalidValueError(value, unit)
	}
	return Size(lo), nil
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

// BytesInt returns size value in bytes as int data type.
// Returned ok is true if size value in bytes is suitable for int data type.
func (s Size) BytesInt() (value int, ok bool) {
	if s > math.MaxInt {
		return 0, false
	}
	return int(s), true
}

var maxUint = Size(math.MaxUint)

// BytesUint returns size value in bytes as uint data type.
// Returned ok is true if size value in bytes is suitable for uint data type.
func (s Size) BytesUint() (value uint, ok bool) {
	if s > maxUint {
		return 0, false
	}
	return uint(s), true
}

// BytesInt32 returns size value in bytes as int32 data type.
// Returned ok is true if size value in bytes is suitable for int32 data type.
func (s Size) BytesInt32() (value int32, ok bool) {
	if s > math.MaxInt32 {
		return 0, false
	}
	return int32(s), true
}

// BytesUint32 returns size value in bytes as uint32 data type.
// Returned ok is true if size value in bytes is suitable for uint32 data type.
func (s Size) BytesUint32() (value uint32, ok bool) {
	if s > math.MaxUint32 {
		return 0, false
	}
	return uint32(s), true
}

// BytesInt64 returns size value in bytes as int64 data type.
// Returned ok is true if size value in bytes is suitable for int64 data type.
func (s Size) BytesInt64() (value int64, ok bool) {
	if s > math.MaxInt64 {
		return 0, false
	}
	return int64(s), true
}

// BytesUint64 returns size value in bytes as uint64 data type.
// Returned ok is true if size value in bytes is suitable for uint64 data type.
func (s Size) BytesUint64() (value uint64, ok bool) {
	return uint64(s), true
}

// BytesFloat32 returns size value in bytes as float32 data type.
// Returned ok is true if size value in bytes is suitable for float32 data type without lost precision.
func (s Size) BytesFloat32() (value float32, ok bool) {
	f := float32(s)
	if Size(f) != s {
		return 0, false
	}
	return f, true
}

// BytesFloat64 returns size value in bytes as float64 data type.
// Returned ok is true if size value in bytes is suitable for float64 data type without lost precision.
func (s Size) BytesFloat64() (value float64, ok bool) {
	f := float64(s)
	if Size(f) != s {
		return 0, false
	}
	return f, true
}

// MarshalText converts size to text.
// If DisableMarshalTextUnit is false, Formatter is used.
// Otherwise, text form in bytes without unit is returned.
//
// See also DisableMarshalTextUnit.
func (s Size) MarshalText() ([]byte, error) {
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

// UnmarshalText using global Parser function.
// DefaultRule affects UnmarshalText behavior.
func (s *Size) UnmarshalText(data []byte) error {
	v, err := Parser(data, DefaultRule&ruleUnmarshalTextMask)
	if err != nil {
		return err
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
		b, err := s.MarshalText()
		if err != nil {
			return nil, err
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
		return err
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
