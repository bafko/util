package test

import (
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
