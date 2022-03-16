// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"go.lstv.dev/util/constraint"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultComparePreRelease(t *testing.T) {
	assert.Equal(t, 0, DefaultComparePreRelease("", ""))
	assert.Equal(t, 1, DefaultComparePreRelease("", "a"))
	assert.Equal(t, -1, DefaultComparePreRelease("a", ""))
	assert.Equal(t, 0, DefaultComparePreRelease("a", "a"))
	assert.Equal(t, -1, DefaultComparePreRelease("a", "aa"))
	assert.Equal(t, 1, DefaultComparePreRelease("aa", "a"))
	assert.Equal(t, -1, DefaultComparePreRelease("aa", "ab"))
	assert.Equal(t, 1, DefaultComparePreRelease("ab", "aa"))
	assert.Equal(t, -1, DefaultComparePreRelease("a", "a1"))
	assert.Equal(t, 1, DefaultComparePreRelease("a1", "a"))
	assert.Equal(t, 0, DefaultComparePreRelease("a1", "a1"))
	assert.Equal(t, 0, DefaultComparePreRelease("a01", "a01"))
	assert.Equal(t, 0, DefaultComparePreRelease("a01", "a1"))
	assert.Equal(t, 0, DefaultComparePreRelease("a1", "a01"))
	assert.Equal(t, -1, DefaultComparePreRelease("a01", "a02"))
	assert.Equal(t, -1, DefaultComparePreRelease("a01", "a2"))
	assert.Equal(t, -1, DefaultComparePreRelease("a1", "a02"))
	assert.Equal(t, 1, DefaultComparePreRelease("a02", "a01"))
	assert.Equal(t, 1, DefaultComparePreRelease("a02", "a1"))
	assert.Equal(t, 1, DefaultComparePreRelease("a2", "a01"))
}

func Test_CompareVersion(t *testing.T) {
	testCompareVersion[string, string](t, CompareVersion[string, string])
}

func testCompareVersion[T1, T2 constraint.ParserInput](t *testing.T, f func(T1, T2) (int, error)) {
	t.Helper()
	c, err := f(T1("0.0.0"), T2("0.0.0"))
	assert.Equal(t, 0, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0"), T2("0.0.0"))
	assert.Equal(t, 1, c)
	assert.NoError(t, err)
	c, err = f(T1("0.0.0"), T2("1.0.0"))
	assert.Equal(t, -1, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0-alfa.1"), T2("1.0.0-alfa.1"))
	assert.Equal(t, 0, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0"), T2("1.0.0-alfa.1"))
	assert.Equal(t, 1, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0-alfa.1"), T2("1.0.0"))
	assert.Equal(t, -1, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0-alfa.2"), T2("1.0.0-alfa.1"))
	assert.Equal(t, 1, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0.0-alfa.1"), T2("1.0.0-alfa.2"))
	assert.Equal(t, -1, c)
	assert.NoError(t, err)
	c, err = f(T1("1.0"), T2("0.0.0"))
	assert.Empty(t, c)
	assert.Error(t, err)
	c, err = f(T1("0.0.0"), T2("1.0"))
	assert.Empty(t, c)
	assert.Error(t, err)
}

func Test_CompareTag(t *testing.T) {
	testCompareTag[string, string](t, CompareTag[string, string])
}

func testCompareTag[T1, T2 constraint.ParserInput](t *testing.T, f func(T1, T2) (int, error)) {
	t.Helper()
	c, err := f(T1("v0.0.0"), T2("v0.0.0"))
	assert.Equal(t, 0, c)
	assert.NoError(t, err)
	c, err = f(T1("v1.0.0"), T2("v0.0.0"))
	assert.Equal(t, 1, c)
	assert.NoError(t, err)
	c, err = f(T1("v0.0.0"), T2("v1.0.0"))
	assert.Equal(t, -1, c)
	assert.NoError(t, err)
	c, err = f(T1("v1.0"), T2("v0.0.0"))
	assert.Empty(t, c)
	assert.Error(t, err)
	c, err = f(T1("v0.0.0"), T2("v1.0"))
	assert.Empty(t, c)
	assert.Error(t, err)
}

func Test_Compare(t *testing.T) {
	testCompareVersion[string, string](t, Compare[string, string])
	testCompareTag[string, string](t, Compare[string, string])
}
