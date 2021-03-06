// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"go.lstv.dev/util/constraint"
)

// Valid checks if passed value is valid roman number.
// If not, error is returned.
//
// See also MaxInputLength.
func Valid[T constraint.ParserInput](input T, r Rule) error {
	const funcName = "Valid"
	empty, err := checkInputLength(funcName, input, r)
	if err != nil {
		return err
	}
	if empty {
		return nil
	}
	if !pattern.Match([]byte(input)) {
		return newNumberFormatError(funcName, input, nil)
	}
	return nil
}
