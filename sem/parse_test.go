// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"go.lstv.dev/util/constraint"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultParser(t *testing.T) {
	v, err := DefaultParser([]byte(`1.2.3`), 0)
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultParser([]byte(`v1.2.3`), 0)
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultParser([]byte(`x`), 0)
	assert.Zero(t, v)
	assert.Error(t, err)
	v, err = DefaultParser([]byte(`1.2.3`), RuleDisableTag)
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v)
	assert.NoError(t, err)
	v, err = DefaultParser([]byte(`v1.2.3`), RuleDisableTag)
	assert.Zero(t, v)
	assert.Error(t, err)
	v, err = DefaultParser([]byte(`x`), RuleDisableTag)
	assert.Zero(t, v)
	assert.Error(t, err)
}

func Test_ParseVersion(t *testing.T) {
	testParseVersion(t, ParseVersion[string])
}

func testParseVersion[T constraint.ParserInput](t *testing.T, f func(T) (Ver, error)) {
	t.Helper()
	valid := map[string]Ver{
		"0.0.0": {
			Major:      0,
			Minor:      0,
			Patch:      0,
			PreRelease: "",
			Build:      "",
		},
		"0.0.1-alpha": {
			Major:      0,
			Minor:      0,
			Patch:      1,
			PreRelease: "alpha",
			Build:      "",
		},
		"0.0.0+abcd": {
			Major:      0,
			Minor:      0,
			Patch:      0,
			PreRelease: "",
			Build:      "abcd",
		},
		"1.0.0": {
			Major:      1,
			Minor:      0,
			Patch:      0,
			PreRelease: "",
			Build:      "",
		},
	}
	invalid := []string{
		"0",
		"1",
		"1.0",
		"1.0.0.0",
	}
	for in, out := range valid {
		v, err := f(T(in))
		assert.Equal(t, out, v)
		assert.NoError(t, err)
	}
	for _, in := range invalid {
		v, err := f(T(in))
		assert.Empty(t, v)
		assert.Error(t, err)
	}
}

func Test_ParseTag(t *testing.T) {
	testParseTag(t, ParseTag[string])
}

func testParseTag[T constraint.ParserInput](t *testing.T, f func(T) (Ver, error)) {
	t.Helper()
	valid := map[string]Ver{
		"v0.0.0": {
			Major:      0,
			Minor:      0,
			Patch:      0,
			PreRelease: "",
			Build:      "",
		},
		"v1.0.0": {
			Major:      1,
			Minor:      0,
			Patch:      0,
			PreRelease: "",
			Build:      "",
		},
	}
	invalid := []string{
		"x",
	}
	for in, out := range valid {
		v, err := f(T(in))
		assert.Equal(t, out, v)
		assert.NoError(t, err)
	}
	for _, in := range invalid {
		v, err := f(T(in))
		assert.Empty(t, v)
		assert.Error(t, err)
	}
}

func Test_Parse(t *testing.T) {
	testParseVersion(t, Parse[string])
	testParseTag(t, Parse[string])
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
	MaxInputLength = 0
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
	MaxInputLength = 4
	assertUnmarshalTextFail(t, `input too long: 5 > 4`, `xxxxx`, 0)
	MaxInputLength = 0
}
