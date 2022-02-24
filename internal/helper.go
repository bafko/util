// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package internal

import (
	"bytes"
	"fmt"
)

// Bprintf formats according to a format specifier and append it to passed buffer.
// It returns passed buffer with added text.
func Bprintf(buf []byte, format string, a ...interface{}) []byte {
	b := bytes.NewBuffer(buf)
	// fmt.Fprintf calls b.Write, which never returns error
	fmt.Fprintf(b, format, a...)
	return b.Bytes()
}
