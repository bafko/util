// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_MarshalBinary(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(5)
	cases := []CaseBinary[mockData[[]byte]]{
		{ // 0
			Data:  []byte(``),
			Value: mockData[[]byte]{[]byte(``)},
		},
		{ // 1
			Constraint: OnlyUnmarshal,
			Data:       []byte(``),
			Value:      mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 2
			Data:  []byte(`abc`),
			Value: mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 3
			Constraint: OnlyMarshal,
			Data:       []byte(`abc`),
			Value:      mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 4
			Data:  []byte(`error`),
			Error: Error(`error`),
			Value: mockData[[]byte]{[]byte(`error`)},
		},
	}
	MarshalBinary(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalBinary_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything)
	m.On("FailNow")
	cases := []CaseBinary[int]{
		{ // 0
			Data:  []byte(``),
			Value: 0,
		},
	}
	MarshalBinary(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalBinary_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	MarshalBinary(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalBinary_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			Before: func(index int, c *CaseBinary[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalBinary(m, cases)
	m.AssertExpectations(t)
}

func Test_MarshalBinary_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			After: func(index int, c *CaseBinary[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	MarshalBinary(m, cases)
	m.AssertExpectations(t)
}

func Test_UnmarshalBinary(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(9)
	cases := []CaseBinary[mockData[[]byte]]{
		{ // 0
			Data:  []byte(``),
			Value: mockData[[]byte]{[]byte(``)},
		},
		{ // 1
			Constraint: OnlyMarshal,
			Data:       []byte(``),
			Value:      mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 2
			Data:  []byte(`abc`),
			Value: mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 3
			Constraint: OnlyUnmarshal,
			Data:       []byte(`abc`),
			Value:      mockData[[]byte]{[]byte(`abc`)},
		},
		{ // 4
			Error: Error(`error`),
			Data:  []byte(`error`),
			Value: mockData[[]byte]{[]byte(`error`)},
		},
	}
	UnmarshalBinary(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalBinary_FailType(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(4)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	m.On("FailNow")
	cases := []CaseBinary[int]{
		{ // 0
			Data:  []byte(``),
			Value: 0,
		},
	}
	UnmarshalBinary(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalBinary_Panic(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			Error: ErrorHasPrefix("panic: failed\n"),
			Value: mockData[string]{mockPanic},
		},
	}
	UnmarshalBinary(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalBinary_Before(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			Before: func(index int, c *CaseBinary[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalBinary(m, cases, nil)
	m.AssertExpectations(t)
}

func Test_UnmarshalBinary_After(t *testing.T) {
	m := &mockT{}
	m.On("Helper").Times(3)
	m.On("Errorf", "\n%s", mock.Anything).Times(1)
	cases := []CaseBinary[mockData[string]]{
		{
			After: func(index int, c *CaseBinary[mockData[string]]) error {
				return errors.New("custom error")
			},
		},
	}
	UnmarshalBinary(m, cases, nil)
	m.AssertExpectations(t)
}
