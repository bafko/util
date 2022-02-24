// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"fmt"
)

// InvalidUnitError represents invalid unit.
// Value of error is invalid unit itself.
type InvalidUnitError string

func newInvalidUnitError(unit string) *InvalidUnitError {
	err := InvalidUnitError(unit)
	return &err
}

// Error returns string representation of error.
func (e *InvalidUnitError) Error() string {
	const message = "invalid unit"
	if e == nil {
		return message
	}
	return fmt.Sprintf("%s %q", message, string(*e))
}

// InvalidValueError represents invalid combination of value and unit.
type InvalidValueError struct {
	Value uint64
	Unit  string
}

func newInvalidValueError(value uint64, unit string) *InvalidValueError {
	return &InvalidValueError{
		Value: value,
		Unit:  unit,
	}
}

// Error returns string representation of error.
func (e *InvalidValueError) Error() string {
	if e.Unit == "" {
		return fmt.Sprintf("value %d without unit is not suitable for uint64", e.Value)
	}
	return fmt.Sprintf("value %d with unit %q is not suitable for uint64", e.Value, e.Unit)
}

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
	err := "unable to parse"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if e.Input == "" {
		return fmt.Sprintf("size.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("size.%s: parsing %q: %s", e.Func, e.Input, err)
}
