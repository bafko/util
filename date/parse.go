// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	// MaxInputLength allows limiting DefaultParser input.
	// Set 0 to disable this setting.
	// ErrInputTooLong is wrapped and used if limit is exceeded.
	// Note: For year > 9999, increment MaxInputLength to > 10.
	MaxInputLength = 10

	// Parser is used by Date.UnmarshalText function.
	Parser = DefaultParser

	pattern = regexp.MustCompile(`^([0-9]{4,9})-?(1[0-2]|0[0-9])-?(3[01]|[0-2][0-9])$`)
)

type (
	// Rule allows configuring Parser behavior.
	Rule int
)

const (
	// RuleDisableBasic disallow basic format (i.e. YYYYMMDD).
	RuleDisableBasic = Rule(1 << iota)
)

// DefaultParser parse Date from input.
//
// See also MaxInputLength.
func DefaultParser(input []byte, r Rule) (date Date, err error) {
	const funcName = "DefaultParser"
	l := len(input)
	if l == 0 {
		return Date{}, newParseError(funcName, "", nil)
	}
	if MaxInputLength != 0 && l > MaxInputLength {
		// do not use input for "input too long" error
		return Date{}, newParseError(funcName, "", fmt.Errorf("%w: %d > %d", ErrInputTooLong, l, MaxInputLength))
	}
	parts := pattern.FindSubmatch(input)
	if len(parts) == 0 {
		return Date{}, newParseError(funcName, string(input), nil)
	}
	if sep2 := input[l-3] == '-'; sep2 || input[l-5] == '-' { // extended format
		if !sep2 || input[l-6] != '-' { // disallow YYYY-MMDD and YYYYMM-YY formats
			return Date{}, newParseError(funcName, string(input), nil)
		}
	} else if r&RuleDisableBasic != 0 {
		return Date{}, newParseError(funcName, string(input), ErrBasicFormatDisabled)
	}
	year, _ := strconv.Atoi(string(parts[1]))
	month, _ := strconv.Atoi(string(parts[2]))
	day, _ := strconv.Atoi(string(parts[3]))
	return New(year, Month(month), day), nil
}
