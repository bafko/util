// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bprintf(t *testing.T) {
	assert.Nil(t, Bprintf(nil, ""))
	assert.Equal(t, []byte(`1`), Bprintf(nil, "%d", 1))
	assert.Equal(t, []byte(`x1`), Bprintf([]byte(`x`), "%d", 1))
}
