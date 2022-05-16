// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// CaseJSON represents one specific test case for MarshalJSON and/or UnmarshalJSON method.
type CaseJSON[T any] struct {
	Constraint Constraint
	Before     func(index int, c *CaseJSON[T]) error
	After      func(index int, c *CaseJSON[T]) error
	Error      AssertErrorFunc
	Data       string
	Value      T
	Custom     any // user-defined custom value for Before and/or After function
}

// MarshalJSON runs all passed test cases of method with same name.
func MarshalJSON[T any](t TestingT, cases []CaseJSON[T]) {
	t.Helper()

	for i, c := range cases {
		if i == 0 {
			if _, ok := any(c.Value).(json.Marshaler); !ok {
				assert.FailNowf(t, "unable to test MarshalJSON", "type %T must implements json.Marshaler", c.Value)
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
		b, err := safeMarshalJSON(any(c.Value).(json.Marshaler))
		if !assert.NoError(t, callForCase(i, &c, c.After), failInfo) {
			continue
		}
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

// UnmarshalJSON runs all passed test cases of method with same name.
func UnmarshalJSON[T any](t TestingT, cases []CaseJSON[T], helper TypeHelper[T]) {
	t.Helper()

	var f func(*T) json.Unmarshaler
	for i, c := range cases {
		if i == 0 {
			if f = castToFunc[T, json.Unmarshaler](c.Value); f == nil {
				assert.FailNowf(t, "unable to test UnmarshalJSON", "type %T must implements json.Unmarshaler", c.Value)
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
		err := safeUnmarshalJSON(f(&v).(json.Unmarshaler), []byte(c.Data))
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

func safeMarshalJSON(m json.Marshaler) (data []byte, err error) {
	defer func() {
		err = panicError(err, recover())
	}()
	return m.MarshalJSON()
}

func safeUnmarshalJSON(u json.Unmarshaler, data []byte) (err error) {
	defer func() {
		err = panicError(err, recover())
	}()
	return u.UnmarshalJSON(data)
}
