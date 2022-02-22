// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseVersion(t *testing.T) {
	testParseVersion(t, ParseVersion)
}

func testParseVersion(t *testing.T, f func([]byte) (Ver, error)) {
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
		v, err := f([]byte(in))
		assert.Equal(t, out, v)
		assert.NoError(t, err)
	}
	for _, in := range invalid {
		v, err := f([]byte(in))
		assert.Empty(t, v)
		assert.Error(t, err)
	}
}

func Test_ParseTag(t *testing.T) {
	testParseTag(t, ParseTag)
}

func testParseTag(t *testing.T, f func([]byte) (Ver, error)) {
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
		v, err := f([]byte(in))
		assert.Equal(t, out, v)
		assert.NoError(t, err)
	}
	for _, in := range invalid {
		v, err := f([]byte(in))
		assert.Empty(t, v)
		assert.Error(t, err)
	}
}

func Test_Parse(t *testing.T) {
	testParseVersion(t, Parse)
	testParseTag(t, Parse)
}
