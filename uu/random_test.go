// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomID(t *testing.T) {
	random = rand.New(rand.NewSource(0))
	id := RandomID()
	assert.Equal(t, `f1f85ff5-85fb-4401-8fad-82097fe9a0e0`, id.String())
}
