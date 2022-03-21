package test

import (
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// CaseJSON represents one specific test case for MarshalJSON and/or UnmarshalJSON method.
type CaseJSON[T any] struct {
	Constraint Constraint
	Error      AssertErrorFunc
	Data       string
	Value      T
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
		b, err := any(c.Value).(json.Marshaler).MarshalJSON()
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

	for i, c := range cases {
		if i == 0 {
			if _, ok := any(&c.Value).(json.Unmarshaler); !ok {
				assert.FailNowf(t, "unable to test UnmarshalJSON", "type *%T must implements json.Unmarshaler", c.Value)
				return
			}
		}

		if !isForUnmarshal(c.Constraint) {
			continue
		}

		failInfo := fmt.Sprintf("case %d failed", i)
		v := helperNew[T](helper, &c.Value)
		err := any(v).(json.Unmarshaler).UnmarshalJSON([]byte(c.Data))
		if c.Error != nil {
			if c.Error(t, err, failInfo) {
				helperAssertEmpty(helper, t, v, failInfo)
			}
		} else {
			if assert.NoError(t, err, failInfo) {
				helperAssertEqual(helper, t, &c.Value, v, failInfo)
			}
		}
	}
}
