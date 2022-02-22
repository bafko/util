// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

const (
	// Byte is unit representing a byte.
	Byte = "B"

	// Kilobyte is unit representing 1000^1 times a byte.
	Kilobyte = "kB"
	// Megabyte is unit representing 1000^2 times a byte.
	Megabyte = "MB"
	// Gigabyte is unit representing 1000^3 times a byte.
	Gigabyte = "GB"
	// Terabyte is unit representing 1000^4 times a byte.
	Terabyte = "TB"
	// Petabyte is unit representing 1000^5 times a byte.
	Petabyte = "PB"
	// Exabyte is unit representing 1000^6 times a byte.
	Exabyte = "EB"
	// Zettabyte is unit representing 1000^7 times a byte.
	// This unit is too big to store in uint64, so only 0 value is allowed.
	Zettabyte = "ZB"
	// Yottabyte is unit representing 1000^8 times a byte.
	// This unit is too big to store in uint64, so only 0 value is allowed.
	Yottabyte = "YB"

	// Kibibyte is unit representing 1024^1 times a byte.
	Kibibyte = "KiB"
	// Mebibyte is unit representing 1024^2 times a byte.
	Mebibyte = "MiB"
	// Gibibyte is unit representing 1024^3 times a byte.
	Gibibyte = "GiB"
	// Tebibyte is unit representing 1024^4 times a byte.
	Tebibyte = "TiB"
	// Pebibyte is unit representing 1024^5 times a byte.
	Pebibyte = "PiB"
	// Exbibyte is unit representing 1024^6 times a byte.
	Exbibyte = "EiB"
	// Zebibyte is unit representing 1024^7 times a byte.
	// This unit is too big to store in uint64, so only 0 value is allowed.
	Zebibyte = "ZiB"
	// Yobibyte is unit representing 1024^8 times a byte.
	// This unit is too big to store in uint64, so only 0 value is allowed.
	Yobibyte = "YiB"
)

var (
	shortenUnits = []string{
		Byte,
		Kibibyte,
		Mebibyte,
		Gibibyte,
		Tebibyte,
		Pebibyte,
	}

	unitToValues = map[string]uint64{
		Byte:     1,
		Kilobyte: 1000,
		Megabyte: 1000 * 1000,
		Gigabyte: 1000 * 1000 * 1000,
		Terabyte: 1000 * 1000 * 1000 * 1000,
		Petabyte: 1000 * 1000 * 1000 * 1000 * 1000,
		Exabyte:  1000 * 1000 * 1000 * 1000 * 1000 * 1000,
		Kibibyte: 1024,
		Mebibyte: 1024 * 1024,
		Gibibyte: 1024 * 1024 * 1024,
		Tebibyte: 1024 * 1024 * 1024 * 1024,
		Pebibyte: 1024 * 1024 * 1024 * 1024 * 1024,
		Exbibyte: 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
	}

	zeroUnits = map[string]struct{}{
		"":        {},
		Byte:      {},
		Kilobyte:  {},
		Megabyte:  {},
		Gigabyte:  {},
		Terabyte:  {},
		Petabyte:  {},
		Exabyte:   {},
		Zettabyte: {},
		Yottabyte: {},
		Kibibyte:  {},
		Mebibyte:  {},
		Gibibyte:  {},
		Tebibyte:  {},
		Pebibyte:  {},
		Exbibyte:  {},
		Zebibyte:  {},
		Yobibyte:  {},
	}
)
