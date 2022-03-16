// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package constraint

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsFloat(t *testing.T) {
	assert.False(t, IsFloat[int]())
	assert.False(t, IsFloat[int8]())
	assert.False(t, IsFloat[int16]())
	assert.False(t, IsFloat[int32]())
	assert.False(t, IsFloat[int64]())
	assert.False(t, IsFloat[uint]())
	assert.False(t, IsFloat[uint8]())
	assert.False(t, IsFloat[uint16]())
	assert.False(t, IsFloat[uint32]())
	assert.False(t, IsFloat[uint64]())
	assert.True(t, IsFloat[float32]())
	assert.True(t, IsFloat[float64]())
}

func Test_IsSigned(t *testing.T) {
	assert.True(t, IsSigned[int]())
	assert.True(t, IsSigned[int8]())
	assert.True(t, IsSigned[int16]())
	assert.True(t, IsSigned[int32]())
	assert.True(t, IsSigned[int64]())
	assert.False(t, IsSigned[uint]())
	assert.False(t, IsSigned[uint8]())
	assert.False(t, IsSigned[uint16]())
	assert.False(t, IsSigned[uint32]())
	assert.False(t, IsSigned[uint64]())
	assert.True(t, IsSigned[float32]())
	assert.True(t, IsSigned[float64]())
}

func Test_SmallestNonzero(t *testing.T) {
	assert.Equal(t, int(1), SmallestNonzero[int]())
	assert.Equal(t, int8(1), SmallestNonzero[int8]())
	assert.Equal(t, int16(1), SmallestNonzero[int16]())
	assert.Equal(t, int32(1), SmallestNonzero[int32]())
	assert.Equal(t, int64(1), SmallestNonzero[int64]())
	assert.Equal(t, uint(1), SmallestNonzero[uint]())
	assert.Equal(t, uint8(1), SmallestNonzero[uint8]())
	assert.Equal(t, uint16(1), SmallestNonzero[uint16]())
	assert.Equal(t, uint32(1), SmallestNonzero[uint32]())
	assert.Equal(t, uint64(1), SmallestNonzero[uint64]())
	assert.Equal(t, float32(math.SmallestNonzeroFloat32), SmallestNonzero[float32]())
	assert.Equal(t, float64(math.SmallestNonzeroFloat64), SmallestNonzero[float64]())
}

func Test_Min(t *testing.T) {
	assert.Equal(t, int(math.MinInt), Min[int]())
	assert.Equal(t, int8(math.MinInt8), Min[int8]())
	assert.Equal(t, int16(math.MinInt16), Min[int16]())
	assert.Equal(t, int32(math.MinInt32), Min[int32]())
	assert.Equal(t, int64(math.MinInt64), Min[int64]())
	assert.Equal(t, uint(0), Min[uint]())
	assert.Equal(t, uint8(0), Min[uint8]())
	assert.Equal(t, uint16(0), Min[uint16]())
	assert.Equal(t, uint32(0), Min[uint32]())
	assert.Equal(t, uint64(0), Min[uint64]())
	assert.Equal(t, float32(-math.MaxFloat32), Min[float32]())
	assert.Equal(t, float64(-math.MaxFloat64), Min[float64]())
}

func Test_Max(t *testing.T) {
	assert.Equal(t, int(math.MaxInt), Max[int]())
	assert.Equal(t, int8(math.MaxInt8), Max[int8]())
	assert.Equal(t, int16(math.MaxInt16), Max[int16]())
	assert.Equal(t, int32(math.MaxInt32), Max[int32]())
	assert.Equal(t, int64(math.MaxInt64), Max[int64]())
	assert.Equal(t, uint(math.MaxUint), Max[uint]())
	assert.Equal(t, uint8(math.MaxUint8), Max[uint8]())
	assert.Equal(t, uint16(math.MaxUint16), Max[uint16]())
	assert.Equal(t, uint32(math.MaxUint32), Max[uint32]())
	assert.Equal(t, uint64(math.MaxUint64), Max[uint64]())
	assert.Equal(t, float32(math.MaxFloat32), Max[float32]())
	assert.Equal(t, float64(math.MaxFloat64), Max[float64]())
}

func Test_SizeBytes(t *testing.T) {
	intSizeBytes := 8
	if math.MaxInt == math.MaxInt32 {
		intSizeBytes = 4
	}
	assert.Equal(t, intSizeBytes, SizeBytes[int]())
	assert.Equal(t, 1, SizeBytes[int8]())
	assert.Equal(t, 2, SizeBytes[int16]())
	assert.Equal(t, 4, SizeBytes[int32]())
	assert.Equal(t, 8, SizeBytes[int64]())
	assert.Equal(t, intSizeBytes, SizeBytes[uint]())
	assert.Equal(t, 1, SizeBytes[uint8]())
	assert.Equal(t, 2, SizeBytes[uint16]())
	assert.Equal(t, 4, SizeBytes[uint32]())
	assert.Equal(t, 8, SizeBytes[uint64]())
	assert.Equal(t, 4, SizeBytes[float32]())
	assert.Equal(t, 8, SizeBytes[float64]())
}

func Test_SizeBits(t *testing.T) {
	intSizeBits := 64
	if math.MaxInt == math.MaxInt32 {
		intSizeBits = 32
	}
	assert.Equal(t, intSizeBits, SizeBits[int]())
	assert.Equal(t, 8, SizeBits[int8]())
	assert.Equal(t, 16, SizeBits[int16]())
	assert.Equal(t, 32, SizeBits[int32]())
	assert.Equal(t, 64, SizeBits[int64]())
	assert.Equal(t, intSizeBits, SizeBits[uint]())
	assert.Equal(t, 8, SizeBits[uint8]())
	assert.Equal(t, 16, SizeBits[uint16]())
	assert.Equal(t, 32, SizeBits[uint32]())
	assert.Equal(t, 64, SizeBits[uint64]())
	assert.Equal(t, 32, SizeBits[float32]())
	assert.Equal(t, 64, SizeBits[float64]())
}
