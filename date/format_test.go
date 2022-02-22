// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultFormatter(t *testing.T, expected string, d Date, f Format) {
	t.Helper()
	b, err := DefaultFormatter(nil, d, f)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
	b, err = DefaultFormatter([]byte("AB"), d, f)
	assert.NoError(t, err)
	assert.Equal(t, "AB"+expected, string(b))
}

func Test_DefaultFormatter(t *testing.T) {
	assertDefaultFormatter(t, `0001-01-01`, Date{}, 0)
	assertDefaultFormatter(t, `2002-08-07`, New(2002, August, 7), 0)
	assertDefaultFormatter(t, `00010101`, Date{}, FormatBasic)
	assertDefaultFormatter(t, `20020807`, New(2002, August, 7), FormatBasic)
}
