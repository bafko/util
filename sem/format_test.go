// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultFormatter(t *testing.T, expected string, v Ver, f Format) {
	t.Helper()
	b, err := DefaultFormatter(nil, v, f)
	assert.Equal(t, expected, string(b))
	assert.NoError(t, err)
	b, err = DefaultFormatter([]byte(`AB`), v, f)
	assert.Equal(t, `AB`+expected, string(b))
	assert.NoError(t, err)
}

func Test_DefaultFormatter(t *testing.T) {
	assertDefaultFormatter(t, `0.0.0`, Ver{}, 0)
	assertDefaultFormatter(t, `v0.0.0`, Ver{}, FormatTag)
	assertDefaultFormatter(t, `1.2.3`, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, 0)
	assertDefaultFormatter(t, `1.2.3-x`, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "",
	}, 0)
	assertDefaultFormatter(t, `1.2.3+y`, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "y",
	}, 0)
	assertDefaultFormatter(t, `1.2.3-x+y`, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "y",
	}, 0)
}
