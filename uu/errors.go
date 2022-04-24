// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"errors"
	"fmt"
	"unicode"

	"go.lstv.dev/util/constraint"
)

var (
	// ErrInputTooLong is wrapped and returned by ID.UnmarshalBinary if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")

	// ErrURNFormatDisabled is wrapped and returned by DefaultParser if RuleDisableURN is present and input is URN form of UUID.
	// Use errors.Is to check if returned error is ErrURNFormatDisabled.
	ErrURNFormatDisabled = errors.New("urn format disabled")
)

// ParseError represents error during uuid parsing.
// Input can be empty, as same as Err.
type ParseError[T constraint.ParserInput] struct {
	Func  string
	Input T
	Err   error
}

func newParseError[T constraint.ParserInput](funcName string, input T, err error) *ParseError[T] {
	return &ParseError[T]{
		Func:  funcName,
		Input: input,
		Err:   err,
	}
}

// Unwrap returns under-laying error if any.
func (e *ParseError[T]) Unwrap() error {
	return e.Err
}

// Error returns string representation of error.
func (e *ParseError[T]) Error() string {
	err := "invalid format"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if len(e.Input) == 0 {
		return fmt.Sprintf("uu.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("uu.%s: %q: %s", e.Func, e.Input, err)
}

// InvalidDigitError represents error returned if invalid digit is found.
type InvalidDigitError byte

// Error returns string representation of error.
func (e InvalidDigitError) Error() string {
	if unicode.IsGraphic(rune(e)) {
		return fmt.Sprintf("invalid digit '%c' (%[1]U)", byte(e))
	}
	return fmt.Sprintf("invalid digit %U", byte(e))
}
