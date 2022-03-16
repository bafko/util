// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"go.lstv.dev/util/constraint"

	"github.com/stretchr/testify/assert"
)

func Test_LatestVersion(t *testing.T) {
	testLatestVersion[string, string](t, LatestVersion[string, string])
}

func testLatestVersion[T1, T2 constraint.ParserInput](t *testing.T, f func(T1, T2) (Ver, error)) {
	t.Helper()
	l, err := f(T1("0.0.0"), T2("0.0.0"))
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("1.0.0"), T2("0.0.0"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("0.0.0"), T2("1.0.0"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("1.0"), T2("0.0.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f(T1("0.0.0"), T2("1.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f(T1("0.0"), T2("1.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
}

func Test_LatestTag(t *testing.T) {
	testLatestTag[string, string](t, LatestTag[string, string])
}

func testLatestTag[T1, T2 constraint.ParserInput](t *testing.T, f func(T1, T2) (Ver, error)) {
	t.Helper()
	l, err := f(T1("v0.0.0"), T2("v0.0.0"))
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("v1.0.0"), T2("v0.0.0"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("v0.0.0"), T2("v1.0.0"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, l)
	assert.NoError(t, err)
	l, err = f(T1("v1.0"), T2("v0.0.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f(T1("v0.0.0"), T2("v1.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
	l, err = f(T1("v0.0"), T2("v1.0"))
	assert.Empty(t, l)
	assert.Error(t, err)
}

func Test_Latest(t *testing.T) {
	testLatestVersion[string, string](t, Latest[string, string])
	testLatestTag[string, string](t, Latest[string, string])
}
