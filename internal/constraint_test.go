// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package internal

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max_unreachable(t *testing.T) {
	assert.Equal(t, int(0), Max[int](Kind("")))
}

func Test_IsFloat_and_IsSigned(t *testing.T) {
	minKind := reflect.Invalid
	maxKind := reflect.UnsafePointer
	for k := minKind; k < maxKind; k++ {
		assert.Equal(t, isFloatKind(k), IsFloat(k))
		assert.Equal(t, isSignedKind(k), IsSigned(k))
	}
}

func Test_SmallestNonzero(t *testing.T) {
	assert.Equal(t, byte(1), SmallestNonzero[byte](reflect.Uint8))
	assert.Equal(t, float32(math.SmallestNonzeroFloat32), SmallestNonzero[float32](reflect.Float32))
	assert.Equal(t, float64(math.SmallestNonzeroFloat64), SmallestNonzero[float64](reflect.Float64))
}

func Test_Min(t *testing.T) {
	assert.Equal(t, int(math.MinInt), Min[int](reflect.Int))
	assert.Equal(t, int8(math.MinInt8), Min[int8](reflect.Int8))
	assert.Equal(t, int16(math.MinInt16), Min[int16](reflect.Int16))
	assert.Equal(t, int32(math.MinInt32), Min[int32](reflect.Int32))
	assert.Equal(t, int64(math.MinInt64), Min[int64](reflect.Int64))
	assert.Equal(t, uint(0), Min[uint](reflect.Uint))
	assert.Equal(t, uint8(0), Min[uint8](reflect.Uint8))
	assert.Equal(t, uint16(0), Min[uint16](reflect.Uint16))
	assert.Equal(t, uint32(0), Min[uint32](reflect.Uint32))
	assert.Equal(t, uint64(0), Min[uint64](reflect.Uint64))
	assert.Equal(t, float32(-math.MaxFloat32), Min[float32](reflect.Float32))
	assert.Equal(t, float64(-math.MaxFloat64), Min[float64](reflect.Float64))
}

func Test_Max(t *testing.T) {
	assert.Equal(t, int(math.MaxInt), Max[int](reflect.Int))
	assert.Equal(t, int8(math.MaxInt8), Max[int8](reflect.Int8))
	assert.Equal(t, int16(math.MaxInt16), Max[int16](reflect.Int16))
	assert.Equal(t, int32(math.MaxInt32), Max[int32](reflect.Int32))
	assert.Equal(t, int64(math.MaxInt64), Max[int64](reflect.Int64))
	assert.Equal(t, uint(math.MaxUint), Max[uint](reflect.Uint))
	assert.Equal(t, uint8(math.MaxUint8), Max[uint8](reflect.Uint8))
	assert.Equal(t, uint16(math.MaxUint16), Max[uint16](reflect.Uint16))
	assert.Equal(t, uint32(math.MaxUint32), Max[uint32](reflect.Uint32))
	assert.Equal(t, uint64(math.MaxUint64), Max[uint64](reflect.Uint64))
	assert.Equal(t, float32(math.MaxFloat32), Max[float32](reflect.Float32))
	assert.Equal(t, float64(math.MaxFloat64), Max[float64](reflect.Float64))
}

func isFloatKind(kind reflect.Kind) bool {
	for _, k := range []reflect.Kind{reflect.Float32, reflect.Float64} {
		if k == kind {
			return true
		}
	}
	return false
}

func isSignedKind(kind reflect.Kind) bool {
	for _, k := range []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64} {
		if k == kind {
			return true
		}
	}
	return false
}
