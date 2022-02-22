// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Valid(t *testing.T) {
	DisableEmptyAsZero = false
	MaxTextLength = 4
	assert.NoError(t, Valid([]byte(nil)))
	assert.NoError(t, Valid([]byte(``)))
	assert.EqualError(t, Valid([]byte(`xxxxx`)), `roman.Valid: input too long (5 > 4)`)
	assert.EqualError(t, Valid([]byte(`yy`)), `roman.Valid: "yy": invalid roman number`)
	assert.NoError(t, Valid([]byte(`ii`)))
	DisableEmptyAsZero = true
	assert.EqualError(t, Valid([]byte(nil)), `roman.Valid: invalid roman number`)
	assert.EqualError(t, Valid([]byte(``)), `roman.Valid: invalid roman number`)
}
