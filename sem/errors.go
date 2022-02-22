// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidPreRelease can be returned by Valid function if pre-release part is not valid.
	ErrInvalidPreRelease = errors.New("invalid pre-release")

	// ErrInvalidBuild can be returned by Valid function if build part is not valid.
	ErrInvalidBuild = errors.New("invalid build")
)

// ParseError represents error during version parsing.
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
	err := "invalid version"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if e.Input == "" {
		return fmt.Sprintf("sem.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("sem.%s: %q: %s", e.Func, e.Input, err)
}
