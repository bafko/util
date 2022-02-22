// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultFormatter(t *testing.T, expected string, s Size, f Format) {
	t.Helper()
	b, err := DefaultFormatter(nil, s, f)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
	b, err = DefaultFormatter([]byte("AB"), s, f)
	assert.NoError(t, err)
	assert.Equal(t, "AB"+expected, string(b))
}

func Test_DefaultFormatter(t *testing.T) {
	assertDefaultFormatter(t, "0B", 0, 0)
	assertDefaultFormatter(t, "11111111B", 11111111, 0)
	assertDefaultFormatter(t, "111111111B", 111111111, 0)
	assertDefaultFormatter(t, "1111111111B", 1111111111, 0)
	assertDefaultFormatter(t, "11111111111B", 11111111111, 0)
	assertDefaultFormatter(t, "0 B", 0, FormatPretty)
	assertDefaultFormatter(t, "11 111 111 B", 11111111, FormatPretty)
	assertDefaultFormatter(t, "111 111 111 B", 111111111, FormatPretty)
	assertDefaultFormatter(t, "1 111 111 111 B", 1111111111, FormatPretty)
	assertDefaultFormatter(t, "11 111 111 111 B", 11111111111, FormatPretty)
}

func Test_appendSeparator(t *testing.T) {
	b := []byte(`10`)
	b = appendSeparator(b, 0)
	assert.Equal(t, []byte(`10`), b)
	b = appendSeparator(b, FormatPretty)
	assert.Equal(t, []byte(`10 `), b)
	b = appendSeparator(b, FormatHTML)
	assert.Equal(t, []byte(`10 `), b)
	b = appendSeparator(b, FormatPretty|FormatHTML)
	assert.Equal(t, []byte(`10 &nbsp;`), b)
}
