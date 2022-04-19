// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ID(t *testing.T) {
	var i ID
	require.Error(t, i.UnmarshalText(nil))
	require.Error(t, i.UnmarshalText([]byte("ed7059f3x8044x4f2ax81aaxb959b33c7777")))
	require.Error(t, i.UnmarshalText([]byte("ed7059fx-8044-4f2a-81aa-b959b33c7777")))
	v := []byte("ed7059f3-8044-4f2a-81aa-b959b33c7777")
	require.NoError(t, i.UnmarshalText(v))
	b, err := i.MarshalText()
	require.NoError(t, err)
	assert.Equal(t, v, b)
}
