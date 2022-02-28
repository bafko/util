// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"fmt"
)

var (
	// MaxInputLength allows limiting Parser input.
	// Set 0 to disable this setting.
	// ErrInputTooLong is wrapped and used if limit is exceeded.
	MaxInputLength = 128

	// Parser is used by Number.UnmarshalText function.
	Parser = DefaultParser
)

type (
	// Rule allows configuring Parser behavior.
	// Available rules are:
	//   RuleDisableEmptyAsZero
	Rule int
)

const (
	// RuleDisableEmptyAsZero force error if empty input is passed to DefaultParser instead of zero as result.
	RuleDisableEmptyAsZero = Rule(1 << iota)
)

// DefaultParser parse roman number from input.
// It returns error if passed value is not valid roman number.
// Parsing is case-insensitive.
// Long and short forms of roman numbers are accepted (e.g. IIII and IV).
//
// See also MaxInputLength.
func DefaultParser(input []byte, r Rule) (Number, error) {
	const funcName = "DefaultParser"
	empty, err := checkInputLength(funcName, input, r)
	if err != nil {
		return 0, err
	}
	if empty {
		return 0, nil
	}
	p := pattern.FindSubmatch(input)
	if len(p) == 0 {
		return 0, newNumberFormatError(funcName, string(input), nil)
	}
	decimal := uint64(len(p[1])) * 1000
	for i, g := range groups {
		decimal += parseGroup(p[i+2], g.Unit, g.Digit5, g.Digit10)
	}
	return Number(decimal), nil
}

func parseGroup(input []byte, unit uint64, digit5, digit10 byte) (decimal uint64) {
	const lowerShift = 'a' - 'A'
	l := uint64(len(input))
	if l == 0 {
		return 0
	}
	if input[0] == digit5 || input[0] == digit5-lowerShift {
		return (4 + l) * unit
	}
	if l == 1 {
		return unit
	}
	if input[1] == digit5 || input[1] == digit5-lowerShift {
		return 4 * unit
	}
	if input[1] == digit10 || input[1] == digit10-lowerShift {
		return 9 * unit
	}
	return l * unit
}

func checkInputLength(funcName string, input []byte, r Rule) (empty bool, err error) {
	l := len(input)
	if l == 0 {
		if r&RuleDisableEmptyAsZero != 0 {
			return false, newNumberFormatError(funcName, "", nil)
		}
		return true, nil
	}
	if MaxInputLength != 0 && l > MaxInputLength {
		// do not use input for "input too long" error
		return false, newNumberFormatError(funcName, "", fmt.Errorf("%w: %d > %d", ErrInputTooLong, l, MaxInputLength))
	}
	return false, nil
}
