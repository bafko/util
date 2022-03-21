// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_Error(t *testing.T) {
	f := Error("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	f(mt, errors.New("abc"), ti)
	mt.AssertExpectations(t)
}

func Test_Error_Fail(t *testing.T) {
	f := Error("abc")
	ti := "testinfo"

	mt := &mockT{}
	mt.On("Helper")
	mt.On("Errorf", "\n%s", mock.Anything)
	f(mt, errors.New("x"), ti)
	mt.AssertExpectations(t)
}
