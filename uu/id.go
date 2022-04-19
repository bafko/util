// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package uu provides type ID to keep, marshal and unmarshal UUID values.
package uu

import (
	"fmt"
)

// ID represents UUID.
type ID struct {
	Higher, Lower uint64
}

// MarshalText converts UUID to text.
func (i ID) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

var (
	starts = []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	digits = map[byte]uint64{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'a': 10,
		'B': 11,
		'b': 11,
		'C': 12,
		'c': 12,
		'D': 13,
		'd': 13,
		'E': 14,
		'e': 14,
		'F': 15,
		'f': 15,
	}
)

// UnmarshalText parse UUID from text.
func (i *ID) UnmarshalText(b []byte) error {
	const length = 36
	if l := len(b); l != length {
		return fmt.Errorf("uu.ID.UnmarshalText: invalid length: expected %d instead of %d", length, l)
	}
	if b[8] != '-' || b[13] != '-' || b[18] != '-' || b[23] != '-' {
		return fmt.Errorf("uu.ID.UnmarshalText: %q: invalid format", b)
	}
	n := [2]uint64{}
	for i, start := range starts {
		for j := 0; j < 2; j++ {
			if v, ok := digits[b[start+j]]; ok {
				x := 124 - (i*2+j)*4
				n[x>>6] |= v << (x & 0x3f)
				continue
			}
			return fmt.Errorf("uu.ID.UnmarshalText: %q: invalid digit %c", b, b[start+j])
		}
	}
	i.Higher = n[1]
	i.Lower = n[0]
	return nil
}

// String formats UUID for string output.
func (i ID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		i.Higher>>32,
		(i.Higher>>16)&0xffff,
		i.Higher&0xffff,
		i.Lower>>48,
		i.Lower&0xffffffffffff,
	)
}
