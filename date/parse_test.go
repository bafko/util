// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultParser(t *testing.T, expected Date, data string, r Rule) {
	t.Helper()
	d, err := DefaultParser([]byte(data), r)
	assert.Equal(t, expected, d)
	assert.NoError(t, err)
}

func assertDefaultParserFail(t *testing.T, error, data string, r Rule) {
	t.Helper()
	d, err := DefaultParser([]byte(data), r)
	assert.Zero(t, d)
	assert.EqualError(t, err, error)
}

func Test_DefaultParser(t *testing.T) {
	MaxTextLength = 0
	assertDefaultParserFail(t, `date.DefaultParser: invalid date`, ``, 0)
	assertDefaultParserFail(t, `date.DefaultParser: "x": invalid date`, `x`, 0)
	assertDefaultParserFail(t, `date.DefaultParser: "2020-0807": invalid date`, `2020-0807`, 0)
	assertDefaultParserFail(t, `date.DefaultParser: "202008-07": invalid date`, `202008-07`, 0)
	assertDefaultParser(t, Date{}, `0001-01-01`, 0)
	assertDefaultParser(t, Date{}, `00010101`, 0)
	assertDefaultParser(t, New(2002, August, 7), `2002-08-07`, 0)
	assertDefaultParser(t, New(2002, August, 7), `20020807`, 0)
	assertDefaultParserFail(t, `date.DefaultParser: "00010101": basic format disabled`, `00010101`, RuleDisableBasic)
	MaxTextLength = 10
	assertDefaultParserFail(t, `date.DefaultParser: input too long: 11 > 10`, `xxxxxxxxxxx`, RuleDisableBasic)
}
