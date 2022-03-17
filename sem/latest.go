// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"go.lstv.dev/util/constraint"
)

// LatestVersion returns the latest version from passed ones.
// If a or b is not valid version, error is returned.
func LatestVersion[T1, T2 constraint.ParserInput](a T1, b T2) (Ver, error) {
	av, err := ParseVersion(a)
	if err != nil {
		return Ver{}, err
	}
	bv, err := ParseVersion(b)
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}

// LatestTag returns the latest tag from passed ones.
// If a or b is not valid tag, error is returned.
func LatestTag[T1, T2 constraint.ParserInput](a T1, b T2) (Ver, error) {
	av, err := ParseTag(a)
	if err != nil {
		return Ver{}, err
	}
	bv, err := ParseTag(b)
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}

// Latest returns the latest version from passed ones.
// If a or b is not valid version or tag, error is returned.
func Latest[T1, T2 constraint.ParserInput](a T1, b T2) (Ver, error) {
	av, err := Parse(a)
	if err != nil {
		return Ver{}, err
	}
	bv, err := Parse(b)
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}
