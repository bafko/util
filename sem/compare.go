// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"fmt"
	"strings"

	"go.lstv.dev/util/constraint"
)

var (
	// ComparePreRelease is used as part of Ver.Compare.
	// If major, minor and patch of compared versions are the same, ComparePreRelease is used.
	ComparePreRelease = DefaultComparePreRelease[string, string]
)

// DefaultComparePreRelease implements https://semver.org/#spec-item-11 rules.
func DefaultComparePreRelease[T1, T2 constraint.ParserInput](a T1, b T2) int {
	la, lb := len(a), len(b)
	if la == 0 {
		if lb == 0 {
			return 0
		}
		return 1
	} else if lb == 0 {
		return -1
	}
	if la > lb {
		return comparePreRelease(b, a)
	}
	return -comparePreRelease(a, b)
}

func comparePreRelease[T1, T2 constraint.ParserInput](shorter T1, longer T2) int {
	s, l := string(shorter), string(longer)
	longerRunes := []rune(l)
	for i, sr := range s {
		if lr := longerRunes[i]; sr != lr {
			return comparePreReleaseSuffix(s[i:], l[i:])
		}
	}
	if len(s) == len(l) {
		return 0
	}
	return 1
}

func comparePreReleaseSuffix(shorter string, longer string) int {
	if digitsOrEmpty.MatchString(shorter) && digitsOrEmpty.MatchString(longer) {
		shorter = strings.TrimLeft(shorter, "0")
		longer = strings.TrimLeft(longer, "0")
	}
	return -strings.Compare(shorter, longer)
}

// CompareVersion compares passed versions.
// Error is returned if passed versions are not valid.
func CompareVersion[T1, T2 constraint.ParserInput](a, b string) (int, error) {
	av, err := ParseVersion(a)
	if err != nil {
		return 0, fmt.Errorf("sem.CompareVersion: %w", err)
	}
	bv, err := ParseVersion(b)
	if err != nil {
		return 0, fmt.Errorf("sem.CompareVersion: %w", err)
	}
	return av.Compare(bv), nil
}

// CompareTag compares passed tag versions.
// Error is returned if passed tag versions are not valid.
func CompareTag[T1, T2 constraint.ParserInput](a T1, b T2) (int, error) {
	av, err := ParseTag(a)
	if err != nil {
		return 0, fmt.Errorf("sem.CompareTag: %w", err)
	}
	bv, err := ParseTag(b)
	if err != nil {
		return 0, fmt.Errorf("sem.CompareTag: %w", err)
	}
	return av.Compare(bv), nil
}

// Compare compares passed tag versions or versions.
// Error is returned if passed tag versions or versions are not valid.
func Compare[T1, T2 constraint.ParserInput](a T1, b T2) (int, error) {
	av, err := Parse(a)
	if err != nil {
		return 0, fmt.Errorf("sem.Compare: %w", err)
	}
	bv, err := Parse(b)
	if err != nil {
		return 0, fmt.Errorf("sem.Compare: %w", err)
	}
	return av.Compare(bv), nil
}
