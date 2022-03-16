// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package internal

import (
	"math"
	"reflect"
)

type ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type uints interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type floats interface {
	~float32 | ~float64
}

type numbers interface {
	ints | uints | floats
}

func Kind(value any) reflect.Kind {
	return reflect.TypeOf(value).Kind()
}

func IsFloat(k reflect.Kind) bool {
	switch k {
	case reflect.Float32, reflect.Float64:
		return true
	default: // reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
		return false
	}
}

func IsSigned(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return true
	default: // reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
		return false
	}
}

func SmallestNonzero[N numbers](k reflect.Kind) N {
	switch k {
	case reflect.Float32:
		return N(any(float32(math.SmallestNonzeroFloat32)).(float32))
	case reflect.Float64:
		return N(any(float64(math.SmallestNonzeroFloat64)).(float64))
	default: // reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
		return 1
	}
}

func Min[N numbers](k reflect.Kind) N {
	switch k {
	case reflect.Int:
		return N(any(int(math.MinInt)).(int))
	case reflect.Int8:
		return N(any(int8(math.MinInt8)).(int8))
	case reflect.Int16:
		return N(any(int16(math.MinInt16)).(int16))
	case reflect.Int32:
		return N(any(int32(math.MinInt32)).(int32))
	case reflect.Int64:
		return N(any(int64(math.MinInt64)).(int64))
	case reflect.Float32:
		return N(any(float32(-math.MaxFloat32)).(float32))
	case reflect.Float64:
		return N(any(float64(-math.MaxFloat64)).(float64))
	default: // reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
		return 0
	}
}

func Max[N numbers](k reflect.Kind) N {
	switch k {
	case reflect.Int:
		return N(any(int(math.MaxInt)).(int))
	case reflect.Int8:
		return N(any(int8(math.MaxInt8)).(int8))
	case reflect.Int16:
		return N(any(int16(math.MaxInt16)).(int16))
	case reflect.Int32:
		return N(any(int32(math.MaxInt32)).(int32))
	case reflect.Int64:
		return N(any(int64(math.MaxInt64)).(int64))
	case reflect.Uint:
		return N(any(uint(math.MaxUint)).(uint))
	case reflect.Uint8:
		return N(any(uint8(math.MaxUint8)).(uint8))
	case reflect.Uint16:
		return N(any(uint16(math.MaxUint16)).(uint16))
	case reflect.Uint32:
		return N(any(uint32(math.MaxUint32)).(uint32))
	case reflect.Uint64:
		return N(any(uint64(math.MaxUint64)).(uint64))
	case reflect.Float32:
		return N(any(float32(math.MaxFloat32)).(float32))
	case reflect.Float64:
		return N(any(float64(math.MaxFloat64)).(float64))
	default: // unreachable
		return 0
	}
}
