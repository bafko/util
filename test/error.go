// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"github.com/stretchr/testify/assert"
)

type AssertErrorFunc func(t TestingT, err error, failInfo string) bool

// Error creates AssertErrorFunc to check if error has passed text.
func Error(text string) AssertErrorFunc {
	return func(t TestingT, err error, failInfo string) bool {
		return assert.EqualError(t, err, text, failInfo)
	}
}
