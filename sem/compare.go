// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"strings"
)

var (
	// ComparePreRelease is used as part of Ver.Compare.
	// If major, minor and patch of compared versions are the same, ComparePreRelease is used.
	ComparePreRelease = DefaultComparePreRelease
)

// DefaultComparePreRelease implements https://semver.org/#spec-item-11 rules.
func DefaultComparePreRelease(a, b string) int {
	if a == "" {
		if b == "" {
			return 0
		}
		return 1
	} else if b == "" {
		return -1
	}
	if len(a) > len(b) {
		return -comparePreRelease(b, a)
	}
	return comparePreRelease(a, b)
}

func comparePreRelease(shorter, longer string) int {
	longerRunes := []rune(longer)
	for i, sr := range shorter {
		if lr := longerRunes[i]; sr != lr {
			return comparePreReleaseSuffix(shorter[i:], longer[i:])
		}
	}
	if len(shorter) == len(longer) {
		return 0
	}
	return 1
}

func comparePreReleaseSuffix(shorter, longer string) int {
	if digitsOrEmpty.MatchString(shorter) && digitsOrEmpty.MatchString(longer) {
		shorter = strings.TrimLeft(shorter, "0")
		longer = strings.TrimLeft(longer, "0")
	}
	return -strings.Compare(shorter, longer)
}

// CompareVersion compares passed versions.
// Error is returned if passed versions are not valid.
func CompareVersion(a, b string) (int, error) {
	av, err := ParseVersion([]byte(a))
	if err != nil {
		return 0, err
	}
	bv, err := ParseVersion([]byte(b))
	if err != nil {
		return 0, err
	}
	return av.Compare(bv), nil
}

// CompareTag compares passed tag versions.
// Error is returned if passed tag versions are not valid.
func CompareTag(a, b string) (int, error) {
	av, err := ParseTag([]byte(a))
	if err != nil {
		return 0, err
	}
	bv, err := ParseTag([]byte(b))
	if err != nil {
		return 0, err
	}
	return av.Compare(bv), nil
}

// Compare compares passed tag versions or versions.
// Error is returned if passed tag versions or versions are not valid.
func Compare(a, b string) (int, error) {
	av, err := Parse([]byte(a))
	if err != nil {
		return 0, err
	}
	bv, err := Parse([]byte(b))
	if err != nil {
		return 0, err
	}
	return av.Compare(bv), nil
}
