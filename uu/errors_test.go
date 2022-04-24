// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newParseError(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, &ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}, newParseError("UnmarshalJSON", "1h", err))
}

func Test_ParseError_Unwrap(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, err, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Unwrap())
}

func Test_ParseError_Error(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, `uu.UnmarshalJSON: invalid format`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   nil,
	}).Error())
	assert.Equal(t, `uu.UnmarshalJSON: parse error`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   err,
	}).Error())
	assert.Equal(t, `uu.UnmarshalJSON: "11": invalid format`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "11",
		Err:   nil,
	}).Error())
	assert.Equal(t, `uu.UnmarshalJSON: "11": parse error`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "11",
		Err:   err,
	}).Error())
}

func Test_InvalidDigitError_Error(t *testing.T) {
	assert.Equal(t, "invalid digit U+0000", InvalidDigitError(0).Error())
	assert.Equal(t, "invalid digit U+000A", InvalidDigitError(10).Error())
	assert.Equal(t, "invalid digit 'x' (U+0078)", InvalidDigitError('x').Error())
}
