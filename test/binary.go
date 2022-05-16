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
	Before     func(index int, c *CaseBinary[T]) error
	After      func(index int, c *CaseBinary[T]) error
	Error      AssertErrorFunc
	Data       []byte
	Value      T
	Custom     any // user-defined custom value for Before and/or After function
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
		if !assert.NoError(t, callForCase(i, &c, c.Before), failInfo) {
			continue
		}
		b, err := safeMarshalBinary(any(c.Value).(encoding.BinaryMarshaler))
		if !assert.NoError(t, callForCase(i, &c, c.After), failInfo) {
			continue
		}
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
		if !assert.NoError(t, callForCase(i, &c, c.Before), failInfo) {
			continue
		}
		v := helperNew[T](helper, c.Value)
		err := safeUnmarshalBinary(f(&v).(encoding.BinaryUnmarshaler), c.Data)
		if !assert.NoError(t, callForCase(i, &c, c.After), failInfo) {
			continue
		}
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

func safeMarshalBinary(m encoding.BinaryMarshaler) (data []byte, err error) {
	defer func() {
		err = panicError(err, recover())
	}()
	return m.MarshalBinary()
}

func safeUnmarshalBinary(u encoding.BinaryUnmarshaler, data []byte) (err error) {
	defer func() {
		err = panicError(err, recover())
	}()
	return u.UnmarshalBinary(data)
}
