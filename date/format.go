// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"go.lstv.dev/util/internal"
)

// Format allows configuring Formatter behavior.
// Available format flags are:
//   FormatBasic
type Format int

const (
	// FormatBasic enforce date format without separators, e.g. YYYYMMDD.
	FormatBasic = Format(1 << iota)
)

var (
	// Formatter is used by Date.MarshalText and other Date converting functions.
	Formatter = DefaultFormatter
)

// DefaultFormatter formats date.
// Default format is ISO 8601 extended format, e.g. YYYY-MM-DD.
// It reacts to Format flags and never returns error.
func DefaultFormatter(buf []byte, d Date, f Format) ([]byte, error) {
	format := `%04d-%02d-%02d`
	if f&FormatBasic != 0 {
		format = `%04d%02d%02d`
	}
	year, month, day := d.Date()
	return internal.Bprintf(buf, format, year, month, day), nil
}
