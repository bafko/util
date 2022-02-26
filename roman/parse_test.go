// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultParser(t *testing.T) {
	MaxTextLength = 14
	valid := map[string]Number{
		"":         0,
		"I":        1,
		"II":       2,
		"III":      3,
		"IIII":     4,
		"IV":       4,
		"V":        5,
		"VI":       6,
		"VII":      7,
		"VIII":     8,
		"VIIII":    9,
		"IX":       9,
		"X":        10,
		"XI":       11,
		"XII":      12,
		"XIII":     13,
		"XIIII":    14,
		"XIV":      14,
		"XV":       15,
		"XVI":      16,
		"XVII":     17,
		"XVIII":    18,
		"XVIIII":   19,
		"XIX":      19,
		"XX":       20,
		"XL":       40,
		"XLI":      41,
		"XLII":     42,
		"XLIII":    43,
		"XLIIII":   44,
		"XLIV":     44,
		"XLV":      45,
		"XC":       90,
		"XCI":      91,
		"XCII":     92,
		"XCIII":    93,
		"XCIIII":   94,
		"XCIV":     94,
		"XCV":      95,
		"CD":       400,
		"CDI":      401,
		"CDV":      405,
		"CDX":      410,
		"CDXLIIII": 444,
		"CDXLIV":   444,
		"CDXLV":    445,
		"CM":       900,
		"CMI":      901,
		"CMV":      905,
		"CMX":      910,
		"CML":      950,
		"CMXLIIII": 944,
		"CMXLIV":   944,
		"CMXLV":    945,
	}
	invalid := []string{
		"IIIII",
		"IVI",
		"VIV",
		"xxxxxxxxxxxxxxx",
	}
	for in, out := range valid {
		actualIn := in
		// test in, "M"+in, "MM"+in...
		for i := Number(0); i < 4; i++ {
			expected := out + (1000 * i)
			actual, err := DefaultParser([]byte(actualIn), 0)
			assert.NoError(t, err)
			assert.Equalf(t, expected, actual, "%s should be %d, not %d", actualIn, expected, actual)
			actualIn = string(thousand) + actualIn
		}
	}
	for _, in := range invalid {
		i, err := DefaultParser([]byte(in), 0)
		assert.Error(t, err)
		assert.Zero(t, i)
	}
	actual, err := DefaultParser([]byte(nil), RuleDisableEmptyAsZero)
	assert.Error(t, err)
	assert.Zero(t, actual)
	actual, err = DefaultParser([]byte(``), RuleDisableEmptyAsZero)
	assert.Error(t, err)
	assert.Zero(t, actual)
}
