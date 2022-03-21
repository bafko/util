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
