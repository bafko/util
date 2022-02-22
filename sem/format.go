// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"strconv"
)

// Format allows configuring Formatter behavior.
// Available format flags are:
//   FormatTag
type Format int

const (
	// FormatTag enforce v prefix.
	FormatTag = Format(1 << iota)
)

var (
	// Formatter is used by Ver.MarshalText and other Ver converting functions.
	Formatter = DefaultFormatter
)

// DefaultFormatter formats version.
// It reacts to Format flags and never returns error.
func DefaultFormatter(buf []byte, v Ver, f Format) ([]byte, error) {
	if f&FormatTag != 0 {
		buf = append(buf, tagPrefix)
	}

	buf = strconv.AppendUint(buf, v.Major, 10)
	buf = append(buf, '.')
	buf = strconv.AppendUint(buf, v.Minor, 10)
	buf = append(buf, '.')
	buf = strconv.AppendUint(buf, v.Patch, 10)

	if v.PreRelease != "" {
		buf = append(buf, '-')
		buf = append(buf, v.PreRelease...)
	}
	if v.Build != "" {
		buf = append(buf, '+')
		buf = append(buf, v.Build...)
	}

	return buf, nil
}
