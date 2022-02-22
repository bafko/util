// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultUnmarshalText(t *testing.T) {
	DisableUnmarshalTag = false
	v, err := DefaultUnmarshalText([]byte(`1.2.3`))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultUnmarshalText([]byte(`v1.2.3`))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultUnmarshalText([]byte(`x`))
	assert.Zero(t, v)
	assert.Error(t, err)
	DisableUnmarshalTag = true
	v, err = DefaultUnmarshalText([]byte(`1.2.3`))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultUnmarshalText([]byte(`v1.2.3`))
	assert.Zero(t, v)
	assert.Error(t, err)
	v, err = DefaultUnmarshalText([]byte(`x`))
	assert.Zero(t, v)
	assert.Error(t, err)
}

func assertUnmarshalText(t *testing.T, expected Ver, input string, f form) {
	t.Helper()
	v, err := unmarshalText("x", []byte(input), f)
	assert.Equal(t, expected, v)
	assert.NoError(t, err)
}

func assertUnmarshalTextFail(t *testing.T, error, input string, f form) {
	t.Helper()
	v, err := unmarshalText("x", []byte(input), f)
	assert.Zero(t, v)
	assert.EqualError(t, err, `sem.x: `+error)
}

func Test_unmarshalText(t *testing.T) {
	MaxTextLength = 0
	assertUnmarshalText(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "y",
	}, `1.2.3-x+y`, formVersion)
	assertUnmarshalText(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "y",
	}, `v1.2.3-x+y`, formTag)
	assertUnmarshalTextFail(t, `invalid version`, ``, 0)
	assertUnmarshalTextFail(t, `"1.2.3": expected tag form`, `1.2.3`, formTag)
	assertUnmarshalTextFail(t, `"v1.2.3": tag form not allowed`, `v1.2.3`, formVersion)
	assertUnmarshalTextFail(t, `"1000000000000000000000000000000.2.3": invalid major`, `1000000000000000000000000000000.2.3`, formVersion)
	assertUnmarshalTextFail(t, `"1.2000000000000000000000000000000.3": invalid minor`, `1.2000000000000000000000000000000.3`, formVersion)
	assertUnmarshalTextFail(t, `"1.2.3000000000000000000000000000000": invalid patch`, `1.2.3000000000000000000000000000000`, formVersion)
	MaxTextLength = 4
	assertUnmarshalTextFail(t, `input too long (5 > 4)`, `xxxxx`, 0)
	MaxTextLength = 0
}
