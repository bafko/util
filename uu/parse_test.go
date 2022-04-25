// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultParser(t *testing.T, expected ID, data string, r Rule) {
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
	MaxInputLength = 0
	assertDefaultParserFail(t, `uu.DefaultParser: invalid format`, ``, 0)
	assertDefaultParserFail(t, `uu.DefaultParser: "x": invalid format`, `x`, 0)
	assertDefaultParserFail(t, `uu.DefaultParser: "0000000000000-0000-0000-000000000000": invalid format`, `0000000000000-0000-0000-000000000000`, 0)
	assertDefaultParserFail(t, `uu.DefaultParser: "00000000-000x-0000-0000-000000000000": invalid digit 'x' (U+0078)`, `00000000-000x-0000-0000-000000000000`, 0)
	assertDefaultParserFail(t, `uu.DefaultParser: "urn:uuid:00000000-000x-0000-0000-000000000000": urn format disabled`, `urn:uuid:00000000-000x-0000-0000-000000000000`, RuleDisableURN)
	assertDefaultParserFail(t, `uu.DefaultParser: "urn:uuid:00000000-0000-0000-0000-000000000000": urn format disabled`, `urn:uuid:00000000-0000-0000-0000-000000000000`, RuleDisableURN)
	assertDefaultParserFail(t, `uu.DefaultParser: "urn:xxxx:0000000000000-0000-0000-000000000000": invalid format`, `urn:xxxx:0000000000000-0000-0000-000000000000`, 0)
	assertDefaultParser(t, ID{}, `00000000-0000-0000-0000-000000000000`, 0)
	assertDefaultParser(t, ID{}, `00000000-0000-0000-0000-000000000000`, RuleDisableURN)
	assertDefaultParser(t, ID{}, `urn:uuid:00000000-0000-0000-0000-000000000000`, 0)
	id := ID{
		Higher: 0xed7059f380444f2a,
		Lower:  0x81aab959b33c7777,
	}
	assertDefaultParser(t, id, `ed7059f3-8044-4f2a-81aa-b959b33c7777`, 0)
	assertDefaultParser(t, id, `ed7059f3-8044-4f2a-81aa-b959b33c7777`, RuleDisableURN)
	assertDefaultParser(t, id, `urn:uuid:ed7059f3-8044-4f2a-81aa-b959b33c7777`, 0)
	assertDefaultParser(t, id, `ed7059f3-8044-4f2A-81aa-b959b33c7777`, 0)
	assertDefaultParser(t, id, `ed7059f3-8044-4f2A-81aa-b959b33c7777`, RuleDisableURN)
	assertDefaultParser(t, id, `urn:uuid:ed7059f3-8044-4f2A-81aa-b959b33c7777`, 0)
	assertDefaultParserFail(t, `uu.DefaultParser: "ed7059f3-8044-4f2A-81aa-b959b33c7777": invalid digit 'A' (U+0041)`, `ed7059f3-8044-4f2A-81aa-b959b33c7777`, RuleDisableUpperCaseDigits)
	assertDefaultParserFail(t, `uu.DefaultParser: "urn:uuid:ed7059f3-8044-4f2A-81aa-b959b33c7777": invalid digit 'A' (U+0041)`, `urn:uuid:ed7059f3-8044-4f2A-81aa-b959b33c7777`, RuleDisableUpperCaseDigits)
	MaxInputLength = 10
	assertDefaultParserFail(t, `uu.DefaultParser: input too long: 11 > 10`, `xxxxxxxxxxx`, 0)
}

func Test_hasURNPrefix(t *testing.T) {
	assert.False(t, hasURNPrefix("         "))
	assert.False(t, hasURNPrefix("arn:uuid:"))
	assert.False(t, hasURNPrefix("uan:uuid:"))
	assert.False(t, hasURNPrefix("Uan:uuid:"))
	assert.False(t, hasURNPrefix("ura:uuid:"))
	assert.False(t, hasURNPrefix("Ura:uuid:"))
	assert.False(t, hasURNPrefix("uRa:uuid:"))
	assert.False(t, hasURNPrefix("URa:uuid:"))
	assert.True(t, hasURNPrefix("urn:uuid:"))
	assert.True(t, hasURNPrefix("Urn:uuid:"))
	assert.True(t, hasURNPrefix("uRn:uuid:"))
	assert.True(t, hasURNPrefix("URn:uuid:"))
	assert.True(t, hasURNPrefix("urN:uuid:"))
	assert.True(t, hasURNPrefix("UrN:uuid:"))
	assert.True(t, hasURNPrefix("uRN:uuid:"))
	assert.True(t, hasURNPrefix("URN:uuid:"))
}
