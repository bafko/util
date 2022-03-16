// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package constraint

import (
	"unsafe"

	"go.lstv.dev/util/internal"
)

// IsFloat returns true if type is float* type.
func IsFloat[N Numbers]() bool {
	return internal.IsFloat(internal.Kind(N(0)))
}

// IsSigned returns true if type is signed integer type or float* type.
func IsSigned[N Numbers]() bool {
	return internal.IsSigned(internal.Kind(N(0)))
}

// SmallestNonzero returns the smallest non-zero value of specified type.
// For integer types returns 1, for float types theirs the smallest non-zero values.
// It always returns positive numbers, so for unsigned integer types, value 1 is returned.
func SmallestNonzero[N Numbers]() N {
	return internal.SmallestNonzero[N](internal.Kind(N(0)))
}

// Min returns the minimum value of specified type.
// For unsigned integer types returns 0, for signed integer types and float types theirs the minimum negative value.
func Min[N Numbers]() N {
	return internal.Min[N](internal.Kind(N(0)))
}

// Max returns the maximum value of specified type.
func Max[N Numbers]() N {
	return internal.Max[N](internal.Kind(N(0)))
}

// SizeBytes returns size in bytes of specified type.
func SizeBytes[N Numbers]() int {
	return int(unsafe.Sizeof(N(0)))
}

// SizeBits returns size in bits of specified type.
func SizeBits[N Numbers]() int {
	return SizeBytes[N]() * 8
}
