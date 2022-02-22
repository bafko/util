// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"fmt"
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
