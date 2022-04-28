// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"encoding"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// CaseBinary represents one specific test case for MarshalBinary and/or UnmarshalBinary method.
type CaseBinary[T any] struct {
	Constraint Constraint
	Error      AssertErrorFunc
	Data       []byte
	Value      T
}

// MarshalBinary runs all passed test cases of method with same name.
func MarshalBinary[T any](t TestingT, cases []CaseBinary[T]) {
	t.Helper()

	for i, c := range cases {
		if i == 0 {
			if _, ok := any(c.Value).(encoding.BinaryMarshaler); !ok {
				assert.FailNowf(t, "unable to test MarshalBinary", "type %T must implements encoding.BinaryMarshaler", c.Value)
				return
			}
		}

		if !isForMarshal(c.Constraint) {
			continue
		}

		failInfo := fmt.Sprintf("case %d failed", i)
		b, err := any(c.Value).(encoding.BinaryMarshaler).MarshalBinary()
		if c.Error != nil {
			if c.Error(t, err, failInfo) {
				assert.Nil(t, b, failInfo)
			}
		} else {
			if assert.NoError(t, err, failInfo) {
				assert.Equal(t, c.Data, b, failInfo)
			}
		}
	}
}

// UnmarshalBinary runs all passed test cases of method with same name.
func UnmarshalBinary[T any](t TestingT, cases []CaseBinary[T], helper TypeHelper[T]) {
	t.Helper()

	var f func(*T) encoding.BinaryUnmarshaler
	for i, c := range cases {
		if i == 0 {
			if f = castToFunc[T, encoding.BinaryUnmarshaler](c.Value); f == nil {
				assert.FailNowf(t, "unable to test UnmarshalBinary", "type %T must implements encoding.BinaryUnmarshaler", c.Value)
				return
			}
		}

		if !isForUnmarshal(c.Constraint) {
			continue
		}

		failInfo := fmt.Sprintf("case %d failed", i)
		v := helperNew[T](helper, c.Value)
		err := f(&v).(encoding.BinaryUnmarshaler).UnmarshalBinary(c.Data)
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
