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

func (m *mockTypeHelperString) New(value string) string {
	args := m.Called(value)
	return args.Get(0).(string)
}

func (m *mockTypeHelperString) AssertEmpty(t TestingT, value string, failInfo string) {
	m.Called(t, value, failInfo)
}

func (m *mockTypeHelperString) AssertEqual(t TestingT, expected, actual string, failInfo string) {
	m.Called(t, expected, actual, failInfo)
}

func Test_helperNew(t *testing.T) {
	empty := ""
	s := "abc"
	assert.Equal(t, empty, helperNew[string](nil, s))

	assert.Equal(t, &empty, helperNew[*string](nil, &s))

	m := &mockTypeHelperString{}
	m.On("New", s).Return(s)
	assert.Equal(t, s, helperNew[string](m, s))
	m.AssertExpectations(t)
}

func Test_helperAssertEmpty(t *testing.T) {
	s := "abc"
	m := &mockTypeHelperString{}
	mt := &mockT{}
	mt.On("Helper")
	m.On("AssertEmpty", mt, s, "case 0 failed")
	helperAssertEmpty[string](m, mt, s, "case 0 failed")
	m.AssertExpectations(t)
}

func Test_helperAssertEqual(t *testing.T) {
	s := "abc"
	m := &mockTypeHelperString{}
	mt := &mockT{}
	mt.On("Helper")
	m.On("AssertEqual", mt, s, s, "case 0 failed")
	helperAssertEqual[string](m, mt, s, s, "case 0 failed")
	m.AssertExpectations(t)
}

type mockXAsValue struct {
	*mock.Mock
}

func (m mockXAsValue) X() {
	m.Called()
}

type mockXAsPointer struct {
	*mock.Mock
}

func (m *mockXAsPointer) X() {
	m.Called()
}

type hasX interface {
	X()
}

func Test_castToFunc(t *testing.T) {
	v := mockXAsValue{
		Mock: &mock.Mock{},
	}
	v.On("X").Once()
	fv := castToFunc[mockXAsValue, hasX](v)
	fv(&v).X()

	p := &mockXAsPointer{
		Mock: &mock.Mock{},
	}
	p.On("X").Once()
	fp := castToFunc[*mockXAsPointer, hasX](p)
	fp(&p).X()

	assert.Nil(t, castToFunc[string, hasX](""))
}
