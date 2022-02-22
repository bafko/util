// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package roman provides type Number to keep and convert roman number to decimal number and vice versa.
package roman

var (
	// DefaultFormat is used at Number.MarshalText and Number.String.
	DefaultFormat = Format(0)
)

// Number represents roman number value.
type Number uint64

// MarshalText converts date to text with Formatter.
//
// See also DefaultFormat.
func (n Number) MarshalText() ([]byte, error) {
	return Formatter(nil, n, DefaultFormat)
}

// UnmarshalText using global UnmarshalText function.
func (n *Number) UnmarshalText(data []byte) error {
	v, err := UnmarshalText(data)
	if err != nil {
		return err
	}
	*n = v
	return nil
}

// String formats roman number for string output.
// If Formatter returns error, String returns same value as DefaultFormatter.
//
// See also DefaultFormat.
func (n Number) String() string {
	b, err := Formatter(nil, n, DefaultFormat)
	if err != nil {
		b, _ = DefaultFormatter(nil, n, DefaultFormat)
		return string(b)
	}
	return string(b)
}
