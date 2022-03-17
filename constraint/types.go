// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package constraint

// ParserInput represents parser input.
type ParserInput interface {
	~[]byte | ~string
}

// Ints represents all int* types and their derived types.
type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Uints represents all uint* types and their derived types.
// Type uintptr is not part of Uints because package aims to numbers with specified value.
type Uints interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Floats represents all float* types and their derived types.
type Floats interface {
	~float32 | ~float64
}

// Numbers represents all number types and their derived types.
// See also Ints, Uints and Floats.
type Numbers interface {
	Ints | Uints | Floats
}
