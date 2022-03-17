// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package roman provides type Number to keep and convert roman number to decimal number and vice versa.
package roman

import (
	"fmt"
)

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

// UnmarshalText using global Parser function.
func (n *Number) UnmarshalText(data []byte) error {
	v, err := Parser(data, 0)
	if err != nil {
		return err
	}
	*n = v
	return nil
}

// Format is implementation for fmt.Formatter.
func (n Number) Format(f fmt.State, verb rune) {
	f.Write(n.format(formatByVerb(verb)))
}

// String formats roman number for string output.
// If Formatter returns error, String returns same value as DefaultFormatter.
//
// See also DefaultFormat.
func (n Number) String() string {
	return string(n.format(DefaultFormat))
}

func (n Number) format(f Format) []byte {
	b, err := Formatter(nil, n, f)
	if err != nil {
		b, _ = DefaultFormatter(nil, n, f)
		return b
	}
	return b
}

func formatByVerb(verb rune) Format {
	switch verb {
	case 'L':
		return FormatLong
	case 'l':
		return FormatLong | FormatLowerCase
	case 'R':
		return 0
	case 'r':
		return FormatLowerCase
	default:
		return DefaultFormat
	}
}
