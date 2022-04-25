// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"errors"
	"fmt"

	"go.lstv.dev/util/constraint"
)

var (
	// ErrInputTooLong is wrapped and returned if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")

	// ErrInvalidPreRelease is wrapped and returned by Valid function if pre-release part is not valid.
	// Use errors.Is to check if returned error is ErrInvalidPreRelease.
	ErrInvalidPreRelease = errors.New("invalid pre-release")

	// ErrInvalidBuild is wrapped and returned by Valid function if build part is not valid.
	// Use errors.Is to check if returned error is ErrInvalidBuild.
	ErrInvalidBuild = errors.New("invalid build")

	// ErrTagFormNotAllowed is wrapped and returned if RuleDisableTag is present and input is tag.
	// Use errors.Is to check if returned error is ErrTagFormNotAllowed.
	ErrTagFormNotAllowed = errors.New("tag form not allowed")

	// ErrExpectedTagForm is wrapped and returned if input is not tag.
	// Use errors.Is to check if returned error is ErrExpectedTagForm.
	ErrExpectedTagForm = errors.New("expected tag form")

	// ErrInvalidMajor is wrapped and returned if input contains invalid major version.
	// Use errors.Is to check if returned error is ErrInvalidMajor.
	ErrInvalidMajor = errors.New("invalid major")

	// ErrInvalidMinor is wrapped and returned if input contains invalid minor version.
	// Use errors.Is to check if returned error is ErrInvalidMinor.
	ErrInvalidMinor = errors.New("invalid minor")

	// ErrInvalidPatch is wrapped and returned if input contains invalid patch version.
	// Use errors.Is to check if returned error is ErrInvalidPatch.
	ErrInvalidPatch = errors.New("invalid patch")
)

// ParseError represents error during version parsing.
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
	err := "invalid version"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if len(e.Input) == 0 {
		return fmt.Sprintf("sem.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("sem.%s: %q: %s", e.Func, e.Input, err)
}
