// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package test contains functions to support advanced testing.
package test

import (
	"reflect"

	"github.com/stretchr/testify/assert"
)

// TestingT is substitute for *testing.T.
type TestingT interface {
	Errorf(format string, args ...any)
	FailNow()
	Helper()
}

// TypeHelper represents object with type manipulation functions.
type TypeHelper[T any] interface {
	New(value T) T
	AssertEmpty(t TestingT, value T, failInfo string)
	AssertEqual(t TestingT, expected, actual T, failInfo string)
}

func helperNew[T any](helper TypeHelper[T], value T) T {
	if helper == nil {
		if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
			return reflect.New(t.Elem()).Interface().(T)
		}
		var newValue T
		return newValue
	}
	return helper.New(value)
}

func helperAssertEmpty[T any](helper TypeHelper[T], t TestingT, value T, failInfo string) {
	t.Helper()
	if helper == nil {
		assert.Empty(t, value, failInfo)
		return
	}
	helper.AssertEmpty(t, value, failInfo)
}

func helperAssertEqual[T any](helper TypeHelper[T], t TestingT, expected, actual T, failInfo string) {
	t.Helper()
	if helper == nil {
		assert.Equal(t, expected, actual, failInfo)
		return
	}
	helper.AssertEqual(t, expected, actual, failInfo)
}

func castToFunc[T, I any](value T) func(*T) I {
	if _, ok := any(value).(I); ok {
		return func(t *T) I {
			return any(*t).(I)
		}
	}
	if _, ok := any(&value).(I); ok {
		return func(t *T) I {
			return any(t).(I)
		}
	}
	return nil
}
