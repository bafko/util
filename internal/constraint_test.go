// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max_unreachable(t *testing.T) {
	assert.Equal(t, int(0), Max[int](Kind("")))
}
