// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"encoding"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// CaseText represents one specific test case for MarshalText and/or UnmarshalText method.
type CaseText[T any] struct {
	Constraint Constraint
	Error      AssertErrorFunc
	Data       string
	Value      T
}

// MarshalText runs all passed test cases of method with same name.
func MarshalText[T any](t TestingT, cases []CaseText[T]) {
	t.Helper()

	for i, c := range cases {
		if i == 0 {
			if _, ok := any(c.Value).(encoding.TextMarshaler); !ok {
				assert.FailNowf(t, "unable to test MarshalText", "type %T must implements encoding.TextMarshaler", c.Value)
				return
			}
		}

		if !isForMarshal(c.Constraint) {
			continue
		}

		failInfo := fmt.Sprintf("case %d failed", i)
		b, err := any(c.Value).(encoding.TextMarshaler).MarshalText()
		if c.Error != nil {
			if c.Error(t, err, failInfo) {
				assert.Nil(t, b, failInfo)
			}
		} else {
			if assert.NoError(t, err, failInfo) {
				assert.Equal(t, c.Data, string(b), failInfo)
			}
		}
	}
}

// UnmarshalText runs all passed test cases of method with same name.
func UnmarshalText[T any](t TestingT, cases []CaseText[T], helper TypeHelper[T]) {
	t.Helper()

	var f func(*T) encoding.TextUnmarshaler
	for i, c := range cases {
		if i == 0 {
			if f = castToFunc[T, encoding.TextUnmarshaler](c.Value); f == nil {
				assert.FailNowf(t, "unable to test UnmarshalText", "type %T must implements encoding.TextUnmarshaler", c.Value)
				return
			}
		}

		if !isForUnmarshal(c.Constraint) {
			continue
		}

		failInfo := fmt.Sprintf("case %d failed", i)
		v := helperNew[T](helper, c.Value)
		err := f(&v).UnmarshalText([]byte(c.Data))
		if c.Error != nil {
			if c.Error(t, err, failInfo) {
				helperAssertEmpty(helper, t, v, failInfo)
			}
		} else {
			if assert.NoError(t, err, failInfo) {
				helperAssertEqual(helper, t, c.Value, v, failInfo)
			}
		}
	}
}
