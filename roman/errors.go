// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"errors"
	"fmt"
)

var (
	// ErrInputTooLong is wrapped and returned by DefaultParser and Valid if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")
)

// NumberFormatError represents error during number parsing.
// Input can be empty, as same as Err.
type NumberFormatError struct {
	Func  string
	Input string
	Err   error
}

func newNumberFormatError(funcName, input string, err error) *NumberFormatError {
	return &NumberFormatError{
		Func:  funcName,
		Input: input,
		Err:   err,
	}
}

// Unwrap returns under-laying error if any.
func (e *NumberFormatError) Unwrap() error {
	return e.Err
}

// Error returns string representation of error.
func (e *NumberFormatError) Error() string {
	err := "invalid roman number"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if e.Input == "" {
		return fmt.Sprintf("roman.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("roman.%s: %q: %s", e.Func, e.Input, err)
}
