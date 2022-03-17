// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"errors"
	"fmt"

	"go.lstv.dev/util/constraint"
)

var (
	// ErrInputTooLong is wrapped and returned by DefaultParser and Valid if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")
)

// NumberFormatError represents error during number parsing.
// Input can be empty, as same as Err.
type NumberFormatError[T constraint.ParserInput] struct {
	Func  string
	Input T
	Err   error
}

func newNumberFormatError[T constraint.ParserInput](funcName string, input T, err error) *NumberFormatError[T] {
	return &NumberFormatError[T]{
		Func:  funcName,
		Input: input,
		Err:   err,
	}
}

// Unwrap returns under-laying error if any.
func (e *NumberFormatError[T]) Unwrap() error {
	return e.Err
}

// Error returns string representation of error.
func (e *NumberFormatError[T]) Error() string {
	err := "invalid roman number"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if len(e.Input) == 0 {
		return fmt.Sprintf("roman.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("roman.%s: %q: %s", e.Func, e.Input, err)
}
