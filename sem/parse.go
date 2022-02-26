// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// MaxTextLength allows limiting Parser input.
	// Set 0 to disable this setting.
	MaxTextLength = 1024

	// Parser is used by Ver.UnmarshalText function.
	Parser = DefaultParser
)

type (
	// Rule allows configuring Parser behavior.
	Rule int
)

const (
	// RuleDisableParseTag disallow tag format.
	// Tag format starts with prefix v.
	RuleDisableParseTag = Rule(1 << iota)
)

// DefaultParser parse Ver from input.
//
// See also MaxTextLength.
func DefaultParser(data []byte, r Rule) (v Ver, err error) {
	const funcName = "DefaultParser"
	f := formVersion
	if r&RuleDisableParseTag == 0 {
		f |= formTag
	}
	return unmarshalText(funcName, data, f)
}

// ParseVersion parses input as version.
// If input is not valid, error is returned.
func ParseVersion(data []byte) (Ver, error) {
	const funcName = "ParseVersion"
	return unmarshalText(funcName, data, formVersion)
}

// ParseTag parses input as tag.
// If input is not valid, error is returned.
func ParseTag(data []byte) (Ver, error) {
	const funcName = "ParseTag"
	return unmarshalText(funcName, data, formTag)
}

// Parse parses input as version or tag.
// If input is not valid, error is returned.
func Parse(data []byte) (Ver, error) {
	const funcName = "Parse"
	return unmarshalText(funcName, data, formVersion|formTag)
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

var (
	errTagFormNotAllowed = errors.New("tag form not allowed")
	errExpectedTagForm   = errors.New("expected tag form")
	errInvalidMajor      = errors.New("invalid major")
	errInvalidMinor      = errors.New("invalid minor")
	errInvalidPatch      = errors.New("invalid patch")
)

func unmarshalText(funcName string, data []byte, f form) (v Ver, err error) {
	l := len(data)
	if l == 0 {
		return Ver{}, newParseError(funcName, "", nil)
	}
	if MaxTextLength != 0 && l > MaxTextLength {
		// do not use input for "input too long" error
		return Ver{}, newParseError(funcName, "", fmt.Errorf("input too long (%d > %d)", l, MaxTextLength))
	}
	if data[0] == tagPrefix {
		if f&formTag == 0 {
			return Ver{}, newParseError(funcName, string(data), errTagFormNotAllowed)
		}
		data = data[1:]
	} else {
		if f&formVersion == 0 {
			return Ver{}, newParseError(funcName, string(data), errExpectedTagForm)
		}
	}
	parts := pattern.FindSubmatch(data)
	if len(parts) == 0 {
		return Ver{}, newParseError(funcName, string(data), nil)
	}
	major, err := strconv.ParseUint(string(parts[1]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, string(data), errInvalidMajor)
	}
	minor, err := strconv.ParseUint(string(parts[2]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, string(data), errInvalidMinor)
	}
	patch, err := strconv.ParseUint(string(parts[3]), 10, 64)
	if err != nil {
		return Ver{}, newParseError(funcName, string(data), errInvalidPatch)
	}
	return Ver{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: string(parts[4]),
		Build:      string(parts[5]),
	}, nil
}
