// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"go.lstv.dev/util/constraint"
)

const (
	// ObjectKeyValue is key "value" for JSON object form.
	ObjectKeyValue = "value"

	// ObjectKeyUnit is key "unit" for JSON object form.
	ObjectKeyUnit = "unit"
)

var (
	// MaxInputLength allows limiting Parser input.
	// Set 0 to disable this setting.
	// ErrInputTooLong is wrapped and used if limit is exceeded.
	MaxInputLength = 128

	// MaxObjectKeys allows limiting UnmarshalJSON object size input.
	// Set 0 to disable this setting.
	// ErrObjectTooBig is wrapped and used if limit is exceeded.
	MaxObjectKeys = 16

	// Parser is used by Size.UnmarshalText and Size.UnmarshalJSON functions.
	Parser = DefaultParser[[]byte]
)

type (
	// Rule allows configuring Parser behavior.
	// Available rules are:
	//   RuleDisableUnit
	//   RuleEnableJSONStringForm
	//   RuleEnableJSONObjectForm
	//   RuleDisallowUnknownKeys
	Rule int
)

const (
	// RuleDisableUnit allows disabling unit specification.
	// If present, Parser returns error if Size is presented with specified unit.
	RuleDisableUnit = Rule(1 << iota)

	// RuleEnableJSONStringForm allows JSON string form.
	// If not present, Parser returns error if Size is presented as JSON string.
	RuleEnableJSONStringForm

	// RuleEnableJSONObjectForm allows JSON object form.
	// If not present, Parser returns error if Size is presented as JSON object.
	RuleEnableJSONObjectForm

	// RuleDisallowUnknownKeys enforce error if JSON object contains other keys than "value" and "unit".
	RuleDisallowUnknownKeys

	ruleIsJSON            = RuleEnableJSONStringForm | RuleEnableJSONObjectForm
	ruleUnmarshalTextMask = RuleDisableUnit

	defaultParserFuncName = "DefaultParser"
)

var (
	// DefaultRule is used by Size.UnmarshalText and Size.UnmarshalJSON converting functions.
	DefaultRule = RuleEnableJSONStringForm | RuleEnableJSONObjectForm
)

// DefaultParser parse Size from input.
// Allowed forms are number, string (with or without specified unit) or JSON object (depends on passed rule value).
//
// For number and string form, spaces before and after data are allowed and ignored.
// Also, spaces/non-breakable-spaces/underscores are allowed and ignored between digits and also as number/unit separator.
//
// JSON object form must contain key "value" with positive or zero number value and "unit" with string value and valid size unit.
// JSON object keys are case-insensitive.
//
// See also MaxInputLength and MaxObjectKeys.
func DefaultParser[T constraint.ParserInput](input T, r Rule) (Size, error) {
	if l := len(input); MaxInputLength != 0 && l > MaxInputLength {
		// do not use input for "input too long" error
		var t T
		return 0, newParseError(defaultParserFuncName, t, fmt.Errorf("%w: %d > %d", ErrInputTooLong, l, MaxInputLength))
	}

	if r&ruleIsJSON != 0 {
		return unmarshalJSON(input, r)
	}

	return unmarshalText(input, r)
}

func unmarshalText[T constraint.ParserInput](input T, r Rule) (Size, error) {
	s := string(input)
	number, unit := prepareNumber(s)

	if number == "" {
		return 0, newParseError(defaultParserFuncName, input, nil)
	}
	value, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, newParseError(defaultParserFuncName, input, err)
	}

	if unit == "" {
		return Size(value), nil
	}
	if r&RuleDisableUnit != 0 {
		return 0, newParseError(defaultParserFuncName, input, ErrUnitDisabled)
	}

	size, err := New(value, unit)
	if err != nil {
		return 0, newParseError(defaultParserFuncName, s, err)
	}
	return size, nil
}

func unmarshalJSON[T constraint.ParserInput](input T, r Rule) (Size, error) {
	d := json.NewDecoder(bytes.NewReader([]byte(input)))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return 0, newParseError(defaultParserFuncName, input, err)
	}
	switch v := t.(type) {
	case json.Delim:
		if v != '{' {
			return 0, newParseError(defaultParserFuncName, input, ErrExpectedObject)
		}
		if r&RuleEnableJSONObjectForm == 0 {
			return 0, newParseError(defaultParserFuncName, input, ErrObjectFormDisabled)
		}
		size, err := unmarshalJSONObject(d, r)
		if err != nil {
			return 0, newParseError(defaultParserFuncName, input, err)
		}
		return size, nil
	case json.Number:
		return unmarshalText([]byte(v), 0)
	case string:
		if r&RuleEnableJSONStringForm == 0 {
			return 0, newParseError(defaultParserFuncName, input, ErrStringFormDisabled)
		}
		return unmarshalText([]byte(v), 0)
	default:
		return 0, newParseError(defaultParserFuncName, input, fmt.Errorf("%w: expected json.Delim, json.Number or string instead of %T", ErrInvalidType, t))
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

func unmarshalJSONObject(d decoder, r Rule) (Size, error) {
	value := (*uint64)(nil)
	unit := (*string)(nil)
keys:
	for i := 0; true; i++ {
		if i > MaxObjectKeys {
			return 0, fmt.Errorf("%w: %d > %d", ErrObjectTooBig, i, MaxObjectKeys)
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
				return 0, ErrDuplicatedValueKey
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
				return 0, ErrDuplicatedUnitKey
			}
			unit, err = decodeUnit(d)
			if err != nil {
				return 0, err
			}
			if !d.More() {
				break keys
			}
		default:
			if r&RuleDisallowUnknownKeys != 0 {
				return 0, fmt.Errorf("%w: %q", ErrUnexpectedKey, key)
			}
			if err := decodeAndSkipNested(d); err != nil {
				return 0, err
			}
		}
	}
	return newOrError(value, unit)
}

func newOrError(value *uint64, unit *string) (Size, error) {
	if value == nil {
		return 0, ErrMissingValueKey
	}
	if unit == nil {
		return 0, ErrMissingUnitKey
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
		return nil, fmt.Errorf("%w: expected json.Number instead of %T for value", ErrInvalidType, t)
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
		return nil, fmt.Errorf("%w: expected string instead of %T for unit", ErrInvalidType, t)
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
