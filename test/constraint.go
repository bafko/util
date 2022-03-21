// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package test

// Constraint allowing limit test cases to specific use.
type Constraint int

const (
	// OnlyMarshal limits test case only for marshal functions.
	OnlyMarshal = Constraint(iota + 1)

	// OnlyUnmarshal limits test case only for unmarshal functions.
	OnlyUnmarshal
)

func isForMarshal(c Constraint) bool {
	return c == 0 || c == OnlyMarshal
}

func isForUnmarshal(c Constraint) bool {
	return c == 0 || c == OnlyUnmarshal
}
