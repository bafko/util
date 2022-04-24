// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"go.lstv.dev/util/internal"
)

const (
	defaultFormat = "%08x-%04x-%04x-%04x-%012x"
	urnFormat     = "urn:uuid:" + defaultFormat
)

// Format allows configuring Formatter behavior.
// Available format flags are:
//   FormatURN
type Format int

const (
	// FormatURN enforce URN form.
	FormatURN = Format(1 << iota)
)

var (
	// Formatter is used by ID.MarshalText and other UUID converting functions.
	Formatter = DefaultFormatter
)

// DefaultFormatter formats UUID.
// It reacts to Format flags and never returns error.
func DefaultFormatter(buf []byte, id ID, f Format) ([]byte, error) {
	format := defaultFormat
	if f&FormatURN != 0 {
		format = urnFormat
	}
	return internal.Bprintf(buf, format,
		id.Higher>>32,
		(id.Higher>>16)&0xffff,
		id.Higher&0xffff,
		id.Lower>>48,
		id.Lower&0xffffffffffff,
	), nil
}
