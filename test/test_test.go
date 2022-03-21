// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTypeHelperString struct {
	mock.Mock
}

func (m *mockTypeHelperString) New(value *string) *string {
	args := m.Called(value)
	return args.Get(0).(*string)
}

func (m *mockTypeHelperString) AssertEmpty(t TestingT, value *string, failInfo string) {
	m.Called(t, value, failInfo)
}

func (m *mockTypeHelperString) AssertEqual(t TestingT, expected, actual *string, failInfo string) {
	m.Called(t, expected, actual, failInfo)
}

func Test_helperNew(t *testing.T) {
	empty := ""
	s := "abc"
	p := &s
	assert.Equal(t, &empty, helperNew[string](nil, p))

	m := &mockTypeHelperString{}
	m.On("New", p).Return(p)
	assert.Equal(t, p, helperNew[string](m, p))
	m.AssertExpectations(t)
}

func Test_helperAssertEmpty(t *testing.T) {
	s := "abc"
	p := &s
	m := &mockTypeHelperString{}
	mt := &mockT{}
	mt.On("Helper")
	m.On("AssertEmpty", mt, p, "case 0 failed").Return(p)
	helperAssertEmpty[string](m, mt, p, "case 0 failed")
	m.AssertExpectations(t)
}

func Test_helperAssertEqual(t *testing.T) {
	s := "abc"
	p := &s
	m := &mockTypeHelperString{}
	mt := &mockT{}
	mt.On("Helper")
	m.On("AssertEqual", mt, p, p, "case 0 failed").Return(p)
	helperAssertEqual[string](m, mt, p, p, "case 0 failed")
	m.AssertExpectations(t)
}
