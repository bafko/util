// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"fmt"
	"strconv"

	"go.lstv.dev/util/constraint"
)

var (
	// MaxInputLength allows limiting Parser input.
	// Set 0 to disable this setting.
	// ErrInputTooLong is wrapped and used if limit is exceeded.
	MaxInputLength = 1024

	// Parser is used by Ver.UnmarshalText function.
	Parser = DefaultParser[[]byte]
)

type (
	// Rule allows configuring Parser behavior.
	// Available rules are:
	//   RuleDisableTag
	Rule int
)

const (
	// RuleDisableTag disallow tag format.
	// Tag format starts with prefix v.
	RuleDisableTag = Rule(1 << iota)
)

// DefaultParser parse Ver from input.
//
// See also MaxInputLength.
func DefaultParser[T constraint.ParserInput](input T, r Rule) (v Ver, err error) {
	const funcName = "DefaultParser"
	f := formVersion
	if r&RuleDisableTag == 0 {
		f |= formTag
	}
	return unmarshalText[T](funcName, input, f)
}

// ParseVersion parses input as version.
// If input is not valid, error is returned.
func ParseVersion[T constraint.ParserInput](input T) (Ver, error) {
	const funcName = "ParseVersion"
	return unmarshalText[T](funcName, input, formVersion)
}

// ParseTag parses input as tag.
// If input is not valid, error is returned.
func ParseTag[T constraint.ParserInput](input T) (Ver, error) {
	const funcName = "ParseTag"
	return unmarshalText[T](funcName, input, formTag)
}

// Parse parses input as version or tag.
// If input is not valid, error is returned.
func Parse[T constraint.ParserInput](input T) (Ver, error) {
	const funcName = "Parse"
	return unmarshalText[T](funcName, input, formVersion|formTag)
}

// form defines which formats are allowed to unmarshal.
// If specified format isn't allowed, unmarshalText returns error.
type form int

const (
	// formVersion for 0.0.0
	formVersion = form(1 << iota)

	// formTag for v0.0.0
	formTag
)

func unmarshalText[T constraint.ParserInput](funcName string, input T, f form) (v Ver, err error) {
	l := len(input)
	if l == 0 {
		return Ver{}, newParseError(funcName, input, nil)
	}
	if MaxInputLength != 0 && l > MaxInputLength {
		// do not use input for "input too long" error
		var t T
		return Ver{}, newParseError(funcName, t, fmt.Errorf("%w: %d > %d", ErrInputTooLong, l, MaxInputLength))
	}
	if input[0] == tagPrefix {
		if f&formTag == 0 {
			return Ver{}, newParseError(funcName, input, ErrTagFormNotAllowed)
		}
		input = input[1:]
	} else {
		if f&formVersion == 0 {
			return Ver{}, newParseError(funcName, input, ErrExpectedTagForm)
		}
	}
	parts := pattern.FindSubmatch([]byte(input))
	if len(parts) == 0 {
		return Ver{}, newParseError(funcName, input, nil)
	}
	major, err := strconv.ParseUint(string(parts[1]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, input, ErrInvalidMajor)
	}
	minor, err := strconv.ParseUint(string(parts[2]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, input, ErrInvalidMinor)
	}
	patch, err := strconv.ParseUint(string(parts[3]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, input, ErrInvalidPatch)
	}
	return Ver{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: string(parts[4]),
		Build:      string(parts[5]),
	}, nil
}
