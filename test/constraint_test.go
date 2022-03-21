package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isForMarshal(t *testing.T) {
	assert.True(t, isForMarshal(0))
	assert.True(t, isForMarshal(OnlyMarshal))
	assert.False(t, isForMarshal(OnlyUnmarshal))
}

func Test_isForUnmarshal(t *testing.T) {
	assert.True(t, isForUnmarshal(0))
	assert.False(t, isForUnmarshal(OnlyMarshal))
	assert.True(t, isForUnmarshal(OnlyUnmarshal))
}
