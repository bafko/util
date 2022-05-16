// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_MarshalJSON(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(5)
	cases := []CaseJSON[mockData[string]]{
		{ // 0
			Data:  ``,
			Value: mockData[string]{``},
		},
		{ // 1
			Constraint: OnlyUnmarshal,
			Data:       ``,
			Value:      mockData[string]{`abc`},
		},
		{ // 2
			Data:  `abc`,
			Value: mockData[string]{`abc`},
		},
		{ // 3
			Constraint: OnlyMarshal,
			Data:       `abc`,
			Value:      mockData[string]{`abc`},
		},
		{ // 4
			Data:  `error`,
			Error: Error(`error`),
			Value: mockData[string]{`error`},
		},
	}
	MarshalJSON(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalJSON_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything)
	m.On("FailNow")
	cases := []CaseJSON[string]{
		{ // 0
			Data:  ``,
			Value: ``,
		},
	}
	MarshalJSON(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalJSON_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	MarshalJSON(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalJSON_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			Before: func(index int, c *CaseJSON[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalJSON(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalJSON_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			After: func(index int, c *CaseJSON[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalJSON(m, cases)
	m.AssertExpectations(t)
}

func Test_UnmarshalJSON(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(9)
	cases := []CaseJSON[mockData[string]]{
		{ // 0
			Data:  ``,
			Value: mockData[string]{``},
		},
		{ // 1
			Constraint: OnlyMarshal,
			Data:       ``,
			Value:      mockData[string]{`abc`},
		},
		{ // 2
			Data:  `abc`,
			Value: mockData[string]{`abc`},
		},
		{ // 3
			Constraint: OnlyUnmarshal,
			Data:       `abc`,
			Value:      mockData[string]{`abc`},
		},
		{ // 4
			Error: Error(`error`),
			Data:  `error`,
			Value: mockData[string]{`error`},
		},
	}
	UnmarshalJSON(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalJSON_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	m.On("FailNow")
	cases := []CaseJSON[string]{
		{ // 0
			Data:  ``,
			Value: ``,
		},
	}
	UnmarshalJSON(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalJSON_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	UnmarshalJSON(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalJSON_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			Before: func(index int, c *CaseJSON[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalJSON(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalJSON_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseJSON[mockData[string]]{
		{
			After: func(index int, c *CaseJSON[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalJSON(m, cases, nil)
	m.AssertExpectations(t)
}
