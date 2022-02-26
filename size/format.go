// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"strconv"
)

var (
	// Formatter is used by Size.MarshalText, Size.MarshalJSON and other Size converting functions.
	Formatter = DefaultFormatter
)

// Format allows configuring Formatter behavior.
// Available format flags are:
//   FormatPretty
//   FormatHTML
type Format int

const (
	// FormatPretty flag forces space between groups of three digits and unit.
	// For example 10000000B is formatted as 10 000 000 B.
	FormatPretty = Format(1 << iota)

	// FormatHTML flag forces format all spaces as "&nbsp;" sequences.
	// DefaultFormatter output can be safely converted to template.HTML if this flag is present.
	FormatHTML
)

// DefaultFormatter for size.
// It reacts to Format flags and never returns error.
func DefaultFormatter(buf []byte, s Size, f Format) ([]byte, error) {
	value, unit := s.Shorten()
	b := []byte(strconv.FormatUint(value, 10))
	offset := 3 - (len(b) % 3)
	for i, digit := range b {
		buf = append(buf, digit)
		if ((i + offset) % 3) == 2 {
			// split to 3-digits long groups
			buf = appendSeparator(buf, f)
		}
	}
	buf = append(buf, unit...)
	return buf, nil
}

func appendSeparator(buf []byte, f Format) []byte {
	if f&FormatPretty == 0 {
		return buf
	}
	if f&FormatHTML == 0 {
		return append(buf, ' ')
	}
	return append(buf, `&nbsp;`...)
}
