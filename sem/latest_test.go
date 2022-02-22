// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LatestVersion(t *testing.T) {
	testLatestVersion(t, LatestVersion)
}

func testLatestVersion(t *testing.T, f func(string, string) (Ver, error)) {
	t.Helper()
	l, err := f("0.0.0", "0.0.0")
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("1.0.0", "0.0.0")
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("0.0.0", "1.0.0")
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("1.0", "0.0.0")
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f("0.0.0", "1.0")
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f("0.0", "1.0")
	assert.Empty(t, l)
	assert.Error(t, err)
}

func Test_LatestTag(t *testing.T) {
	testLatestTag(t, LatestTag)
}

func testLatestTag(t *testing.T, f func(string, string) (Ver, error)) {
	t.Helper()
	l, err := f("v0.0.0", "v0.0.0")
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("v1.0.0", "v0.0.0")
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("v0.0.0", "v1.0.0")
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f("v1.0", "v0.0.0")
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f("v0.0.0", "v1.0")
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f("v0.0", "v1.0")
	assert.Empty(t, l)
	assert.Error(t, err)
}

func Test_Latest(t *testing.T) {
	testLatestVersion(t, Latest)
	testLatestTag(t, Latest)
}
