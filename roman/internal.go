// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"regexp"
)

const thousand = 'M'

var (
	pattern = regexp.MustCompile(`(?i)^(M*)(D?C{0,4}|CD|CM)(L?X{0,4}|XL|XC)(V?I{0,4}|IV|IX)$`)

	hundreds = []string{
		"",
		"C",
		"CC",
		"CCC",
		"CD",
		"D",
		"DC",
		"DCC",
		"DCCC",
		"CM",
	}

	tens = []string{
		"",
		"X",
		"XX",
		"XXX",
		"XL",
		"L",
		"LX",
		"LXX",
		"LXXX",
		"XC",
	}

	units = []string{
		"",
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
	}

	groups = []struct {
		Unit    uint64
		Digit5  byte
		Digit10 byte
	}{
		{
			Unit:    100,
			Digit5:  'D',
			Digit10: 'M',
		},
		{
			Unit:    10,
			Digit5:  'L',
			Digit10: 'C',
		},
		{
			Unit:    1,
			Digit5:  'V',
			Digit10: 'X',
		},
	}
)
