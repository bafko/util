// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newNumberFormatError(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, &NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}, newNumberFormatError("UnmarshalJSON", "1h", err))
}

func Test_NumberFormatError_Unwrap(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, err, (&NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Unwrap())
}

func Test_NumberFormatError_Error(t *testing.T) {
	err := errors.New("parse error")
	assert.Equal(t, `roman.UnmarshalJSON: invalid roman number`, (&NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   nil,
	}).Error())
	assert.Equal(t, `roman.UnmarshalJSON: parse error`, (&NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "",
		Err:   err,
	}).Error())
	assert.Equal(t, `roman.UnmarshalJSON: "1h": invalid roman number`, (&NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   nil,
	}).Error())
	assert.Equal(t, `roman.UnmarshalJSON: "1h": parse error`, (&NumberFormatError[string]{
		Func:  "UnmarshalJSON",
		Input: "1h",
		Err:   err,
	}).Error())
}
