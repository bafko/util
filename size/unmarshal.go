// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// ObjectKeyValue is key "value" for JSON object form.
	ObjectKeyValue = "value"

	// ObjectKeyUnit is key "unit" for JSON object form.
	ObjectKeyUnit = "unit"
)

var (
	// MaxTextLength allows limiting UnmarshalText input.
	// Set 0 to disable this setting.
	MaxTextLength = 128

	// MaxObjectKeys allows limiting UnmarshalJSON object size input.
	// Set 0 to disable this setting.
	MaxObjectKeys = 16

	// DisableUnmarshalTextUnit allows disabling unit specification at DefaultUnmarshalText.
	// If true, DefaultUnmarshalText returns error if Size is presented with specified unit.
	DisableUnmarshalTextUnit = false

	// DisableUnmarshalJSONStringForm allows disabling string form at DefaultUnmarshalJSON.
	// If true, DefaultUnmarshalJSON returns error if Size is presented as JSON string.
	DisableUnmarshalJSONStringForm = false

	// DisableUnmarshalJSONObjectForm allows disabling object form at DefaultUnmarshalJSON.
	// If true, DefaultUnmarshalJSON returns error if Size is presented as JSON object.
	DisableUnmarshalJSONObjectForm = false

	// UnmarshalText is used by Size.UnmarshalText and DefaultUnmarshalJSON functions.
	UnmarshalText = DefaultUnmarshalText

	// UnmarshalJSON is used by Size.UnmarshalJSON function.
	UnmarshalJSON = DefaultUnmarshalJSON
)

// DefaultUnmarshalText parse Size from text.
// Allowed forms are number or string (with or without specified unit).
// Spaces before and after data are allowed and ignored.
// Also, spaces/non-breakable-spaces/underscores are allowed and ignored between digits and also as number/unit separator.
//
// See also MaxTextLength and DisableUnmarshalTextUnit.
func DefaultUnmarshalText(data []byte) (Size, error) {
	const funcName = "DefaultUnmarshalText"

	if l := len(data); MaxTextLength != 0 && l > MaxTextLength {
		// do not use input for "input too long" error
		return 0, newParseError(funcName, "", fmt.Errorf("input too long (%d > %d)", l, MaxTextLength))
	}

	s := string(data)
	number, unit := prepareNumber(s)

	if number == "" {
		return 0, newParseError(funcName, s, nil)
	}
	value, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, newParseError(funcName, s, err)
	}

	if unit == "" {
		return Size(value), nil
	}
	if DisableUnmarshalTextUnit {
		return 0, newParseError(funcName, s, errors.New("unit disabled"))
	}

	size, err := New(value, unit)
	if err != nil {
		return 0, newParseError(funcName, s, err)
	}
	return size, nil
}

// DefaultUnmarshalJSON parse Size from JSON.
// Allowed forms are number, string (with or without specified unit) or object.
//
// Strings are parsed with DefaultUnmarshalText.
//
// Object must contain key "value" with positive or zero number value and "unit" with string value and valid size unit.
// Object keys are case-insensitive.
//
// See also MaxObjectKeys, DisableUnmarshalJSONStringForm and DisableUnmarshalJSONObjectForm.
func DefaultUnmarshalJSON(data []byte) (Size, error) {
	const funcName = "DefaultUnmarshalJSON"

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return 0, newParseError(funcName, string(data), err)
	}
	switch v := t.(type) {
	case json.Delim:
		if v != '{' {
			return 0, newParseError(funcName, string(data), errors.New("expected number, string or object"))
		}
		if DisableUnmarshalJSONObjectForm {
			return 0, newParseError(funcName, string(data), errors.New("object form disabled"))
		}
		size, err := unmarshalJSONObject(d)
		if err != nil {
			return 0, newParseError(funcName, string(data), err)
		}
		return size, nil
	case json.Number:
		return UnmarshalText([]byte(v))
	case string:
		if DisableUnmarshalJSONStringForm {
			return 0, newParseError(funcName, string(data), errors.New("string form disabled"))
		}
		return UnmarshalText([]byte(v))
	default:
		return 0, newParseError(funcName, string(data), fmt.Errorf("unexpected type %T", t))
	}
}

func prepareNumber(input string) (number, unit string) {
	const (
		sp   = ' '
		nbsp = 0xA0
	)
	n := strings.Builder{}
	for i, r := range input {
		if r == sp {
			continue
		}
		if n.Len() > 0 {
			if r == '_' || r == nbsp {
				continue
			}
		}
		if r >= '0' && r <= '9' {
			n.WriteRune(r)
			continue
		}
		// unit remains
		return n.String(), strings.TrimSuffix(input[i:], string(sp))
	}
	return n.String(), ""
}

type decoder interface {
	Token() (json.Token, error)
	More() bool
}

func unmarshalJSONObject(d decoder) (Size, error) {
	value := (*uint64)(nil)
	unit := (*string)(nil)
keys:
	for i := 0; true; i++ {
		if i > MaxObjectKeys {
			return 0, fmt.Errorf("object too big (%d > %d)", i, MaxObjectKeys)
		}
		if !d.More() {
			break keys
		}
		t, err := d.Token()
		if err != nil {
			return 0, err
		}
		// string is guaranteed by encoding/json package
		key := t.(string)
		switch strings.ToLower(key) {
		case ObjectKeyValue:
			if value != nil {
				return 0, errors.New("duplicated value key")
			}
			value, err = decodeValue(d)
			if err != nil {
				return 0, err
			}
			if !d.More() {
				break keys
			}
		case ObjectKeyUnit:
			if unit != nil {
				return 0, errors.New("duplicated unit key")
			}
			unit, err = decodeUnit(d)
			if err != nil {
				return 0, err
			}
			if !d.More() {
				break keys
			}
		default:
			if err := decodeAndSkipNested(d); err != nil {
				return 0, err
			}
		}
	}
	return newOrError(value, unit)
}

func newOrError(value *uint64, unit *string) (Size, error) {
	if value == nil {
		return 0, errors.New("missing value key")
	}
	if unit == nil {
		return 0, errors.New("missing unit key")
	}
	return New(*value, *unit)
}

func decodeValue(d decoder) (*uint64, error) {
	t, err := d.Token()
	if err != nil {
		return nil, err
	}
	n, ok := t.(json.Number)
	if !ok {
		return nil, fmt.Errorf("expected type json.Number instead of %T for value", t)
	}
	u, err := strconv.ParseUint(n.String(), 10, 64)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func decodeUnit(d decoder) (*string, error) {
	t, err := d.Token()
	if err != nil {
		return nil, err
	}
	s, ok := t.(string)
	if !ok {
		return nil, fmt.Errorf("expected type string instead of %T for unit", t)
	}
	return &s, nil
}

func decodeAndSkipNested(d decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch t.(type) {
	case json.Delim:
		// open delim is guaranteed by json.Decoder
		for depth := 1; ; {
			t2, err := d.Token()
			if err != nil {
				return err
			}
			if v, ok := t2.(json.Delim); ok {
				if v == '{' || v == '[' {
					depth++
				} else {
					depth--
				}
			}
			if depth == 0 {
				break
			}
		}
		return nil
	default: // no-op
		return nil
	}
}
