// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
)

// AssertErrorFunc represents function to assert error.
type AssertErrorFunc func(t TestingT, err error, failInfo string) bool

// AnyError is AssertErrorFunc to check if any error was passed.
var AnyError AssertErrorFunc = func(t TestingT, err error, failInfo string) bool {
	return assert.Error(t, err, failInfo)
}

// Error creates AssertErrorFunc to check if error has passed text.
func Error(text string) AssertErrorFunc {
	return func(t TestingT, err error, failInfo string) bool {
		return assert.EqualError(t, err, text, failInfo)
	}
}

// ErrorHasPrefix creates AssertErrorFunc to check if error has passed text prefix.
func ErrorHasPrefix(prefix string) AssertErrorFunc {
	return func(t TestingT, err error, failInfo string) bool {
		if assert.Error(t, err) {
			return assert.True(t, strings.HasPrefix(err.Error(), prefix), failInfo)
		}
		return false
	}
}

// ErrorHasSuffix creates AssertErrorFunc to check if error has passed text suffix.
func ErrorHasSuffix(suffix string) AssertErrorFunc {
	return func(t TestingT, err error, failInfo string) bool {
		if assert.Error(t, err) {
			return assert.True(t, strings.HasSuffix(err.Error(), suffix), failInfo)
		}
		return false
	}
}

// ErrorMatch creates AssertErrorFunc to check if error text match passed regexp.
func ErrorMatch(regexpPattern string) AssertErrorFunc {
	return func(t TestingT, err error, failInfo string) bool {
		if assert.Error(t, err) {
			ok, err := regexp.MatchString(regexpPattern, err.Error())
			if ok {
				return true
			}
			assert.NoError(t, err, failInfo)
			return false
		}
		return false
	}
}
