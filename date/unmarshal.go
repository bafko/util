// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	// MaxTextLength allows limiting UnmarshalText input.
	// Set 0 to disable this setting.
	// Note: For year > 9999, increment MaxTextLength to > 10.
	MaxTextLength = 10

	// DisableUnmarshalBasic disallow basic format at DefaultUnmarshalText.
	DisableUnmarshalBasic = false

	// UnmarshalText is used by Date.UnmarshalText function.
	UnmarshalText = DefaultUnmarshalText

	errBasicFormatDisabled = errors.New("basic format disabled")
	pattern                = regexp.MustCompile(`^([0-9]{4,9})-?(1[0-2]|0[0-9])-?(3[01]|[0-2][0-9])$`)
)

// DefaultUnmarshalText parse Date from text.
//
// See also MaxTextLength and DisableUnmarshalBasic.
func DefaultUnmarshalText(data []byte) (date Date, err error) {
	const funcName = "DefaultUnmarshalText"
	l := len(data)
	if l == 0 {
		return Date{}, newParseError(funcName, "", nil)
	}
	if MaxTextLength != 0 && l > MaxTextLength {
		// do not use input for "input too long" error
		return Date{}, newParseError(funcName, "", fmt.Errorf("input too long (%d > %d)", l, MaxTextLength))
	}
	parts := pattern.FindSubmatch(data)
	if len(parts) == 0 {
		return Date{}, newParseError(funcName, string(data), nil)
	}
	if sep2 := data[l-3] == '-'; sep2 || data[l-5] == '-' { // extended format
		if !sep2 || data[l-6] != '-' { // disallow YYYY-MMDD and YYYYMM-YY formats
			return Date{}, newParseError(funcName, string(data), nil)
		}
	} else if DisableUnmarshalBasic {
		return Date{}, newParseError(funcName, string(data), errBasicFormatDisabled)
	}
	year, _ := strconv.Atoi(string(parts[1]))
	month, _ := strconv.Atoi(string(parts[2]))
	day, _ := strconv.Atoi(string(parts[3]))
	return New(year, Month(month), day), nil
}
