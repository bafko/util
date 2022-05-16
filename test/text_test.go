// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_MarshalText(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(5)
	cases := []CaseText[mockData[string]]{
		{
			Data:  ``,
			Value: mockData[string]{``},
		},
		{
			Constraint: OnlyUnmarshal,
			Data:       ``,
			Value:      mockData[string]{`abc`},
		},
		{
			Data:  `abc`,
			Value: mockData[string]{`abc`},
		},
		{
			Constraint: OnlyMarshal,
			Data:       `abc`,
			Value:      mockData[string]{`abc`},
		},
		{
			Error: Error(`error`),
			Data:  `error`,
			Value: mockData[string]{`error`},
		},
	}
	MarshalText(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalText_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything)
	m.On("FailNow")
	cases := []CaseText[string]{
		{
			Data:  ``,
			Value: ``,
		},
	}
	MarshalText(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalText_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(1)
	cases := []CaseText[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	MarshalText(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalText_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseText[mockData[string]]{
		{
			Before: func(index int, c *CaseText[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalText(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalText_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseText[mockData[string]]{
		{
			After: func(index int, c *CaseText[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalText(m, cases)
	m.AssertExpectations(t)
}

func Test_UnmarshalText(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(9)
	cases := []CaseText[mockData[string]]{
		{
			Data:  ``,
			Value: mockData[string]{``},
		},
		{
			Constraint: OnlyMarshal,
			Data:       ``,
			Value:      mockData[string]{`abc`},
		},
		{
			Data:  `abc`,
			Value: mockData[string]{`abc`},
		},
		{
			Constraint: OnlyUnmarshal,
			Data:       `abc`,
			Value:      mockData[string]{`abc`},
		},
		{
			Error: Error(`error`),
			Data:  `error`,
			Value: mockData[string]{`error`},
		},
	}
	UnmarshalText(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalText_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	m.On("FailNow")
	cases := []CaseText[string]{
		{
			Data:  ``,
			Value: ``,
		},
	}
	UnmarshalText(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalText_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseText[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	UnmarshalText(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalText_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseText[mockData[string]]{
		{
			Before: func(index int, c *CaseText[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalText(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalText_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseText[mockData[string]]{
		{
			After: func(index int, c *CaseText[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalText(m, cases, nil)
	m.AssertExpectations(t)
}
