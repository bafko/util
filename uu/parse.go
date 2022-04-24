// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"fmt"

	"go.lstv.dev/util/constraint"
)

var starts = []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}

var (
	// MaxInputLength allows limiting DefaultParser input.
	// Set 0 to disable this setting.
	// ErrInputTooLong is wrapped and used if limit is exceeded.
	// Note: For year > 9999, increment MaxInputLength to > 45.
	MaxInputLength = 45

	// Parser is used by ID.UnmarshalText function.
	Parser = DefaultParser[[]byte]
)

type (
	// Rule allows configuring Parser behavior.
	// Available rules are:
	//   RuleDisableURN
	//   RuleDisableUpperCaseDigits
	Rule int
)

const (
	// RuleDisableURN disallow URN form.
	RuleDisableURN = Rule(1 << iota)

	// RuleDisableUpperCaseDigits disallow upper-case digits.
	RuleDisableUpperCaseDigits
)

// DefaultParser parse UUID from input.
//
// See also MaxInputLength.
func DefaultParser[T constraint.ParserInput](input T, r Rule) (id ID, err error) {
	const funcName = "DefaultParser"
	l := len(input)
	if MaxInputLength != 0 && l > MaxInputLength {
		// do not use input for "input too long" error
		var t T
		return ID{}, newParseError(funcName, t, fmt.Errorf("%w: %d > %d", ErrInputTooLong, l, MaxInputLength))
	}
	offset := 0
	switch l {
	case IDLength:
		// no-op
	case IDLength + len(URNPrefix):
		if r&RuleDisableURN != 0 {
			return ID{}, newParseError(funcName, input, ErrURNFormatDisabled)
		}
		if !hasURNPrefix(input) {
			return ID{}, newParseError(funcName, input, nil)
		}
		// move offset after URN prefix
		offset = len(URNPrefix)
	default:
		return ID{}, newParseError(funcName, input, nil)
	}
	if input[offset+8] != '-' || input[offset+13] != '-' || input[offset+18] != '-' || input[offset+23] != '-' {
		return ID{}, newParseError(funcName, input, nil)
	}
	n := [2]uint64{}
	allowUpperCase := r&RuleDisableUpperCaseDigits == 0
	for i, start := range starts {
		for j := 0; j < 2; j++ {
			if v, ok := parseDigit(input[offset+start+j], allowUpperCase); ok {
				x := 124 - (i*2+j)*4
				n[x>>6] |= v << (x & 0x3f)
				continue
			}
			return ID{}, newParseError(funcName, input, InvalidDigitError(input[offset+start+j]))
		}
	}
	return ID{
		Higher: n[1],
		Lower:  n[0],
	}, nil
}

func hasURNPrefix[T constraint.ParserInput](input T) bool {
	if input[0] != 'u' && input[0] != 'U' {
		return false
	}
	if input[1] != 'r' && input[1] != 'R' {
		return false
	}
	if input[2] != 'n' && input[2] != 'N' {
		return false
	}
	for i, r := range []byte(":uuid:") {
		if input[i+3] != r {
			return false
		}
	}
	return true
}

func parseDigit(digit byte, allowUpperCase bool) (uint64, bool) {
	if digit >= '0' && digit <= '9' {
		return uint64(digit - '0'), true
	}
	if digit >= 'a' && digit <= 'f' {
		return uint64(digit - ('a' - 10)), true
	}
	if allowUpperCase && digit >= 'A' && digit <= 'F' {
		return uint64(digit - ('A' - 10)), true
	}
	return 0, false
}
