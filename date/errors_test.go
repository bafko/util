// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

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
	assert.Equal(t, `date.UnmarshalJSON: invalid date`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   nil,
	}).Error())
	assert.Equal(t, `date.UnmarshalJSON: parse error`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   err,
	}).Error())
	assert.Equal(t, `date.UnmarshalJSON: "1h": invalid date`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   nil,
	}).Error())
	assert.Equal(t, `date.UnmarshalJSON: "1h": parse error`, (&ParseError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Error())
}
