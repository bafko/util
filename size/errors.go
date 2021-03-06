// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"errors"
	"fmt"

	"go.lstv.dev/util/constraint"
)

var (
	// ErrInputTooLong is wrapped and returned by DefaultParser if passed input is too long.
	// Use errors.Is to check if returned error is ErrInputTooLong.
	ErrInputTooLong = errors.New("input too long")

	// ErrObjectTooBig is wrapped and returned by DefaultParser if passed JSON object has too many keys.
	// Use errors.Is to check if returned error is ErrObjectTooBig.
	ErrObjectTooBig = errors.New("object too big")

	// ErrInvalidType is wrapped and returned by DefaultParser if invalid type is occurred.
	// Use errors.Is to check if returned error is ErrInvalidType.
	ErrInvalidType = errors.New("invalid type")

	// ErrUnitDisabled is wrapped and returned by DefaultParser if RuleDisableUnit is present and input contains unit.
	// Use errors.Is to check if returned error is ErrUnitDisabled.
	ErrUnitDisabled = errors.New("unit disabled")

	// ErrExpectedObject is wrapped and returned by DefaultParser
	// if RuleEnableJSONObjectForm is present and input contains invalid JSON kind.
	// Use errors.Is to check if returned error is ErrExpectedObject.
	ErrExpectedObject = errors.New("expected object")

	// ErrObjectFormDisabled is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm is not present and input contains JSON object.
	// Use errors.Is to check if returned error is ErrObjectFormDisabled.
	ErrObjectFormDisabled = errors.New("object form disabled")

	// ErrStringFormDisabled is wrapped and returned by DefaultParser if RuleEnableJSONStringForm is not present and input contains JSON string.
	// Use errors.Is to check if returned error is ErrStringFormDisabled.
	ErrStringFormDisabled = errors.New("string form disabled")

	// ErrMissingValueKey is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm is present and input contains JSON object without "value" key.
	// Use errors.Is to check if returned error is ErrMissingValueKey.
	ErrMissingValueKey = errors.New("missing value key")

	// ErrMissingUnitKey is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm is present and input contains JSON object without "unit" key.
	// Use errors.Is to check if returned error is ErrMissingUnitKey.
	ErrMissingUnitKey = errors.New("missing unit key")

	// ErrUnexpectedKey is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm and RuleDisallowUnknownKeys is present and input contains JSON object with other than "value" or "unit" keys.
	// Use errors.Is to check if returned error is ErrUnexpectedKey.
	ErrUnexpectedKey = errors.New("unexpected key")

	// ErrDuplicatedValueKey is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm is present and input contains JSON object with duplicated "value" key.
	// Use errors.Is to check if returned error is ErrDuplicatedValueKey.
	ErrDuplicatedValueKey = errors.New("duplicated value key")

	// ErrDuplicatedUnitKey is wrapped and returned by DefaultParser if RuleEnableJSONObjectForm is present and input contains JSON object with duplicated "unit" key.
	// Use errors.Is to check if returned error is ErrDuplicatedUnitKey.
	ErrDuplicatedUnitKey = errors.New("duplicated unit key")
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
type InvalidValueError[N constraint.Numbers] struct {
	Value N
	Unit  string
}

func newInvalidValueError[N constraint.Numbers](value N, unit string) *InvalidValueError[N] {
	return &InvalidValueError[N]{
		Value: value,
		Unit:  unit,
	}
}

// Error returns string representation of error.
func (e *InvalidValueError[N]) Error() string {
	if e.Unit == "" {
		return fmt.Sprintf("value %v without unit is not suitable for uint64", e.Value)
	}
	return fmt.Sprintf("value %v with unit %q is not suitable for uint64", e.Value, e.Unit)
}

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
	err := "unable to parse"
	if e.Err != nil {
		err = e.Err.Error()
	}
	if len(e.Input) == 0 {
		return fmt.Sprintf("size.%s: %s", e.Func, err)
	}
	return fmt.Sprintf("size.%s: parsing %q: %s", e.Func, e.Input, err)
}
