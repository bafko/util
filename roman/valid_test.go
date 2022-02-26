// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Valid(t *testing.T) {
	MaxTextLength = 4
	assert.NoError(t, Valid([]byte(nil), 0))
	assert.NoError(t, Valid([]byte(``), 0))
	assert.EqualError(t, Valid([]byte(`xxxxx`), 0), `roman.Valid: input too long (5 > 4)`)
	assert.EqualError(t, Valid([]byte(`yy`), 0), `roman.Valid: "yy": invalid roman number`)
	assert.NoError(t, Valid([]byte(`ii`), 0))
	assert.EqualError(t, Valid([]byte(nil), RuleDisableEmptyAsZero), `roman.Valid: invalid roman number`)
	assert.EqualError(t, Valid([]byte(``), RuleDisableEmptyAsZero), `roman.Valid: invalid roman number`)
}
