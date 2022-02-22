// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultFormatter(t *testing.T) {
	data := map[Number]string{
		0:    "",
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		9:    "IX",
		19:   "XIX",
		40:   "XL",
		44:   "XLIV",
		90:   "XC",
		94:   "XCIV",
		400:  "CD",
		900:  "CM",
		1000: "M",
		3001: "MMMI",
	}
	for in, out := range data {
		b, err := DefaultFormatter(nil, in, 0)
		assert.Equal(t, out, string(b))
		assert.NoError(t, err)
		b, err = DefaultFormatter([]byte(`AB`), in, 0)
		assert.Equal(t, `AB`+out, string(b))
		assert.NoError(t, err)
	}
}

func Test_toHundreds(t *testing.T) {
	assert.Equal(t, `CD`, toHundreds(4, 0))
	assert.Equal(t, `CCCC`, toHundreds(4, FormatLong400))
	assert.Equal(t, `CM`, toHundreds(9, 0))
	assert.Equal(t, `DCCCC`, toHundreds(9, FormatLong900))
}

func Test_toTens(t *testing.T) {
	assert.Equal(t, `XL`, toTens(4, 0))
	assert.Equal(t, `XXXX`, toTens(4, FormatLong40))
	assert.Equal(t, `XC`, toTens(9, 0))
	assert.Equal(t, `LXXXX`, toTens(9, FormatLong90))
}

func Test_toUnits(t *testing.T) {
	assert.Equal(t, `IV`, toUnits(4, 0))
	assert.Equal(t, `IIII`, toUnits(4, FormatLong4))
	assert.Equal(t, `IX`, toUnits(9, 0))
	assert.Equal(t, `VIIII`, toUnits(9, FormatLong9))
}
