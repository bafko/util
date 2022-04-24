// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultFormatter(t *testing.T, expected string, id ID, f Format) {
	t.Helper()
	b, err := DefaultFormatter(nil, id, f)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
	b, err = DefaultFormatter([]byte("AB"), id, f)
	assert.NoError(t, err)
	assert.Equal(t, "AB"+expected, string(b))
}

func Test_DefaultFormatter(t *testing.T) {
	assertDefaultFormatter(t, `00000000-0000-0000-0000-000000000000`, ID{}, 0)
	assertDefaultFormatter(t, `ed7059f3-8044-4f2a-81aa-b959b33c7777`, ID{
		Higher: 0xed7059f380444f2a,
		Lower:  0x81aab959b33c7777,
	}, 0)
	assertDefaultFormatter(t, `urn:uuid:00000000-0000-0000-0000-000000000000`, ID{}, FormatURN)
	assertDefaultFormatter(t, `urn:uuid:ed7059f3-8044-4f2a-81aa-b959b33c7777`, ID{
		Higher: 0xed7059f380444f2a,
		Lower:  0x81aab959b33c7777,
	}, FormatURN)
}
