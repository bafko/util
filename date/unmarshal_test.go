// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultUnmarshalText(t *testing.T, expected Date, data string) {
	t.Helper()
	d, err := DefaultUnmarshalText([]byte(data))
	assert.Equal(t, expected, d)
	assert.NoError(t, err)
}

func assertDefaultUnmarshalTextFail(t *testing.T, error, data string) {
	t.Helper()
	d, err := DefaultUnmarshalText([]byte(data))
	assert.Zero(t, d)
	assert.EqualError(t, err, error)
}

func Test_DefaultUnmarshalText(t *testing.T) {
	MaxTextLength = 0
	DisableUnmarshalBasic = false
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: invalid date`, ``)
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: "x": invalid date`, `x`)
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: "2020-0807": invalid date`, `2020-0807`)
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: "202008-07": invalid date`, `202008-07`)
	assertDefaultUnmarshalText(t, Date{}, `0001-01-01`)
	assertDefaultUnmarshalText(t, Date{}, `00010101`)
	assertDefaultUnmarshalText(t, New(2002, August, 7), `2002-08-07`)
	assertDefaultUnmarshalText(t, New(2002, August, 7), `20020807`)
	DisableUnmarshalBasic = true
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: "00010101": basic format disabled`, `00010101`)
	MaxTextLength = 10
	assertDefaultUnmarshalTextFail(t, `date.DefaultUnmarshalText: input too long (11 > 10)`, `xxxxxxxxxxx`)
}
