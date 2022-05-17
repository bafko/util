// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_AnyError(t *testing.T) {
	ti := "testinfo"

	mt := &mockT{}
	assert.True(t, AnyError(mt, errors.New("abc"), ti))
	mt.AssertExpectations(t)
}

func Test_AnyError_Fail(t *testing.T) {
	ti := "testinfo"

	mt := &mockT{}
	assert.True(t, AnyError(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}

func Test_Error(t *testing.T) {
	f := Error("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	assert.True(t, f(mt, errors.New("abc"), ti))
	mt.AssertExpectations(t)
}

func Test_Error_Fail(t *testing.T) {
	f := Error("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasPrefix(t *testing.T) {
	f := ErrorHasPrefix("abc")
	ti := "testinfo"

	mt := &mockT{}
	assert.True(t, f(mt, errors.New("abc"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasPrefix_Nil(t *testing.T) {
	f := ErrorHasPrefix("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, nil, ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasPrefix_Fail(t *testing.T) {
	f := ErrorHasPrefix("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasSuffix(t *testing.T) {
	f := ErrorHasSuffix("abc")
	ti := "testinfo"

	mt := &mockT{}
	assert.True(t, f(mt, errors.New("abc"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasSuffix_Nil(t *testing.T) {
	f := ErrorHasSuffix("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, nil, ti))
	mt.AssertExpectations(t)
}

func Test_ErrorHasSuffix_Fail(t *testing.T) {
	f := ErrorHasSuffix("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorMatch(t *testing.T) {
	f := ErrorMatch("ab.")
	ti := "testinfo"

	mt := &mockT{}
	assert.True(t, f(mt, errors.New("abc"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorMatch_Nil(t *testing.T) {
	f := ErrorMatch("ab.")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, nil, ti))
	mt.AssertExpectations(t)
}

func Test_ErrorMatch_Fail(t *testing.T) {
	f := ErrorMatch("ab.")
	ti := "testinfo"

	mt := &mockT{}
	assert.False(t, f(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}

func Test_ErrorMatch_InvalidRegexpPattern(t *testing.T) {
	f := ErrorMatch("ab(.")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	assert.False(t, f(mt, errors.New("x"), ti))
	mt.AssertExpectations(t)
}
