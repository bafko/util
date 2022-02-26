// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"errors"
	"fmt"
)

var (
	// ErrInputTooLong is wrapped and returned by Date.UnmarshalBinary if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")

	// ErrInvalidLength is wrapped and returned by Date.UnmarshalBinary if passed input has invalid length.
	// Use errors.Is to check if returned error is ErrInvalidLength.
	ErrInvalidLength = errors.New("invalid length")

	// ErrUnsupportedVersion is wrapped and returned by Date.UnmarshalBinary if passed input has unsupported version.
	// Use errors.Is to check if returned error is ErrUnsupportedVersion.
	ErrUnsupportedVersion = errors.New("unsupported version")

	// ErrInvalidType is wrapped and returned by Date.Scan if passed type is invalid.
	// Use errors.Is to check if returned error is ErrInvalidType.
	ErrInvalidType = errors.New("invalid type")

	// ErrBasicFormatDisabled is wrapped and returned by DefaultParser if RuleDisableBasic is present and input is basic format date.
	// Use errors.Is to check if returned error is ErrBasicFormatDisabled.
	ErrBasicFormatDisabled = errors.New("basic format disabled")

	// ErrInvalidFromOrTo is wrapped and returned by FilterFromTo if passed from or to is invalid.
	// Use errors.Is to check if returned error is ErrInvalidFromOrTo.
	ErrInvalidFromOrTo = errors.New("invalid from or to")
)

// ParseError represents error during date parsing.
// Input can be empty, as same as Err.
type ParseError struct {
	Func  string
	Input string
	Err   error
}

func newParseError(funcName, input string, err error) *ParseError {
	return &ParseError{
		Func:  funcName,
		Input: input,
		Err:   err,
	}
}

// Unwrap returns under-laying error if any.
func (e *ParseError) Unwrap() error {
	return e.Err
}

// Error returns string representation of error.
func (e *ParseError) Error() string {
	err := "invalid date"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if e.Input == "" {
		return fmt.Sprintf("date.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("date.%s: %q: %s", e.Func, e.Input, err)
}
