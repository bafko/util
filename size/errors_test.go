// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newInvalidUnitError(t *testing.T) {
	unit := "h"
	expected := InvalidUnitError(unit)
	assert.Equal(t, &expected, newInvalidUnitError(unit))
}

func Test_InvalidUnitError_Error(t *testing.T) {
	assert.Equal(t, `invalid unit`, (*InvalidUnitError)(nil).Error())
	err := InvalidUnitError("h")
	assert.Equal(t, `invalid unit "h"`, err.Error())
}

func Test_newInvalidValueError(t *testing.T) {
	assert.Equal(t, &InvalidValueError[uint64]{
		Value: 1,
		Unit:  Zebibyte,
	}, newInvalidValueError(uint64(1), Zebibyte))
}

func Test_InvalidValueError_Error(t *testing.T) {
	assert.Equal(t, `value 1 without unit is not suitable for uint64`, (&InvalidValueError[uint64]{
		Value: 1,
		Unit:  "",
	}).Error())
	assert.Equal(t, `value 1 with unit "ZiB" is not suitable for uint64`, (&InvalidValueError[uint64]{
		Value: 1,
		Unit:  Zebibyte,
	}).Error())
}

func Test_newParseError(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, &ParseError{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}, newParseError("UnmarshalJSON", "1h", err))
}

func Test_ParseError_Unwrap(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, err, (&ParseError{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Unwrap())
}

func Test_ParseError_Error(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, `size.UnmarshalJSON: unable to parse`, (&ParseError{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   nil,
	}).Error())
	assert.Equal(t, `size.UnmarshalJSON: parse error`, (&ParseError{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   err,
	}).Error())
	assert.Equal(t, `size.UnmarshalJSON: parsing "1h": unable to parse`, (&ParseError{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   nil,
	}).Error())
	assert.Equal(t, `size.UnmarshalJSON: parsing "1h": parse error`, (&ParseError{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Error())
}
