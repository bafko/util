// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"bytes"
	"math/bits"
)

var (
	// Formatter is used by Number.MarshalText and other Number converting functions.
	Formatter = DefaultFormatter
)

// Format allows configuring Formatter behavior.
// Available format flags are:
//   FormatLong
//   FormatLong4
//   FormatLong40
//   FormatLong400
//   FormatLong4x
//   FormatLong9
//   FormatLong90
//   FormatLong900
//   FormatLong9x
type Format int

const (
	// FormatLong4 disable short form of 4 (e.g. IV).
	//   IV -> IIII
	FormatLong4 = Format(1 << iota)

	// FormatLong40 disable short form of 40 (e.g. XL).
	//   XL -> XXXX
	FormatLong40

	// FormatLong400 disable short form of 400 (e.g. CD).
	//   CD -> CCCC
	FormatLong400

	// FormatLong9 disable short form of 9 (e.g. IX).
	//   IX -> VIIII
	FormatLong9

	// FormatLong90 disable short form of 9 (e.g. XC).
	//   XC -> LXXXX
	FormatLong90

	// FormatLong900 disable short form of 9 (e.g. CM).
	//   CM -> DCCCC
	FormatLong900

	// FormatLong4x is combination of FormatLong4,  FormatLong40 and FormatLong400.
	FormatLong4x = FormatLong4 | FormatLong40 | FormatLong400

	// FormatLong9x is combination of FormatLong9,  FormatLong90 and FormatLong900.
	FormatLong9x = FormatLong9 | FormatLong90 | FormatLong900

	// FormatLong is combination of FormatLong4x and FormatLong9x.
	FormatLong = FormatLong4x | FormatLong9x
)

// DefaultFormatter converts decimal number to roman number.
// Zero is represented as empty string.
// It reacts to Format flags and never returns error.
func DefaultFormatter(buf []byte, n Number, f Format) ([]byte, error) {
	if n == 0 {
		return buf, nil
	}
	r, i := bits.Div64(0, uint64(n), 1000)
	b := bytes.NewBuffer(buf)
	for j := uint64(0); j < r; j++ {
		b.WriteByte(thousand)
	}
	r, i = bits.Div64(0, i, 100)
	b.WriteString(toHundreds(r, f))
	r, i = bits.Div64(0, i, 10)
	b.WriteString(toTens(r, f))
	b.WriteString(toUnits(i, f))
	return b.Bytes(), nil
}

func toHundreds(value uint64, f Format) string {
	if value == 4 && f&FormatLong400 != 0 {
		return "CCCC"
	}
	if value == 9 && f&FormatLong900 != 0 {
		return "DCCCC"
	}
	return hundreds[value]
}

func toTens(value uint64, f Format) string {
	if value == 4 && f&FormatLong40 != 0 {
		return "XXXX"
	}
	if value == 9 && f&FormatLong90 != 0 {
		return "LXXXX"
	}
	return tens[value]
}

func toUnits(value uint64, f Format) string {
	if value == 4 && f&FormatLong4 != 0 {
		return "IIII"
	}
	if value == 9 && f&FormatLong9 != 0 {
		return "VIIII"
	}
	return units[value]
}
