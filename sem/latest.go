// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

// LatestVersion returns the latest version from passed ones.
// If a or b is not valid version, error is returned.
func LatestVersion(a, b string) (Ver, error) {
	av, err := ParseVersion([]byte(a))
	if err != nil {
		return Ver{}, err
	}
	bv, err := ParseVersion([]byte(b))
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}

// LatestTag returns the latest tag from passed ones.
// If a or b is not valid tag, error is returned.
func LatestTag(a, b string) (Ver, error) {
	av, err := ParseTag([]byte(a))
	if err != nil {
		return Ver{}, err
	}
	bv, err := ParseTag([]byte(b))
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}

// Latest returns the latest version from passed ones.
// If a or b is not valid version or tag, error is returned.
func Latest(a, b string) (Ver, error) {
	av, err := Parse([]byte(a))
	if err != nil {
		return Ver{}, err
	}
	bv, err := Parse([]byte(b))
	if err != nil {
		return Ver{}, err
	}
	return av.Latest(bv), nil
}
