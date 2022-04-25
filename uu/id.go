// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package uu provides type ID to keep, marshal and unmarshal UUID values.
// See also RFC 4122.
package uu

import (
	"fmt"
)

const (
	// URNPrefix represents UUID URN prefix.
	URNPrefix = "urn:uuid:"

	// IDLength represents UUID length.
	IDLength = 36
)

// ID represents UUID.
type ID struct {
	Higher, Lower uint64
}

// Version returns UUID version.
// Returned value has range 0-15.
func (i ID) Version() int {
	return int((i.Higher >> 12) & 0xf)
}

// Variant returns UUID variant.
// Returned value has range 0-3 (count of variant bits).
func (i ID) Variant() int {
	if i.Lower&0x8000000000000000 == 0 {
		return 0
	}
	if i.Lower&0x4000000000000000 == 0 {
		return 1
	}
	if i.Lower&0x2000000000000000 == 0 {
		return 2
	}
	return 3
}

// URN returns UUID in URN form.
func (i ID) URN() string {
	b, _ := DefaultFormatter([]byte("urn:uuid:"), i, 0)
	return string(b)
}

// MarshalText converts UUID to text.
func (i ID) MarshalText() ([]byte, error) {
	b, err := Formatter(nil, i, 0)
	if err != nil {
		return nil, fmt.Errorf("uu.ID.MarshalText: %w", err)
	}
	return b, nil
}

// UnmarshalText parse UUID from text.
func (i *ID) UnmarshalText(data []byte) error {
	id, err := Parser(data, 0)
	if err != nil {
		return fmt.Errorf("uu.ID.UnmarshalText: %w", err)
	}
	*i = id
	return nil
}

// Format is implementation for fmt.Formatter.
//
//   ┌ Verb ┬ Format ───┬ Example ────────────────────────────────────────┐
//   │ %s   │ Format(0) │ "00000000-0000-0000-0000-000000000000"          │
//   │ %u   │ FormatURN │ "urn:uuid:00000000-0000-0000-0000-000000000000" │
func (i ID) Format(f fmt.State, verb rune) {
	f.Write(i.format(formatByVerb(verb)))
}

// String formats UUID for string output.
func (i ID) String() string {
	return string(i.format(0))
}

func (i ID) format(f Format) []byte {
	b, err := Formatter(nil, i, f)
	if err != nil {
		b, _ = DefaultFormatter(nil, i, f)
	}
	return b
}

func formatByVerb(verb rune) Format {
	switch verb {
	case 'u':
		return FormatURN
	default:
		return 0
	}
}
