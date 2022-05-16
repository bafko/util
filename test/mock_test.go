// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

const mockPanic = "PANIC"

type mockT struct {
	mock.Mock
}

func (m *mockT) Helper() {
	m.Called()
}

func (m *mockT) Errorf(format string, args ...any) {
	m.Called(format, args)
}

func (m *mockT) FailNow() {
	m.Called()
}

var errMockData = errors.New("error")

type mockData[T []byte | string] struct {
	data T
}

func (m mockData[T]) MarshalBinary() ([]byte, error) {
	if m.isError() {
		return nil, errMockData
	}
	return []byte(m.data), nil
}

func (m *mockData[T]) UnmarshalBinary(data []byte) error {
	if isError(data) {
		return errMockData
	}
	m.data = T(data)
	return nil
}

func (m mockData[T]) MarshalText() ([]byte, error) {
	if m.isError() {
		return nil, errMockData
	}
	return []byte(m.data), nil
}

func (m *mockData[T]) UnmarshalText(data []byte) error {
	if isError(data) {
		return errMockData
	}
	m.data = T(data)
	return nil
}

func (m mockData[T]) MarshalJSON() ([]byte, error) {
	if m.isError() {
		return nil, errMockData
	}
	return []byte(m.data), nil
}

func (m *mockData[T]) UnmarshalJSON(data []byte) error {
	if isError(data) {
		return errMockData
	}
	m.data = T(data)
	return nil
}

func (m mockData[T]) isError() bool {
	return isError([]byte(m.data))
}

func isError(data []byte) bool {
	if string(data) == mockPanic {
		panic("failed")
	}
	return string(data) == "error"
}
