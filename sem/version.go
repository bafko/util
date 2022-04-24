// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package sem provides type Ver to keep, marshal and unmarshal version values.
// See also: https://semver.org/
package sem

import (
	"fmt"
	"math/bits"
)

// Ver represents version consist of major, minor, patch, pre-release and build components.
// Values major, minor and patch are represented as uint64.
// Pre-release and build are strings, and they are omitted in string form if empty.
// These two components also can be invalid, so Valid method should be used to check, if current value is valid.
type Ver struct {
	Major, Minor, Patch uint64
	PreRelease, Build   string
}

// New creates version with specified major, minor, patch, pre-release and build values.
// It should be called like:
//   New(1, 0, 0)
//   New(1, 0, 0, "alfa")
//   New(1, 0, 0, "alfa", "0123456789")
//   New(1, 0, 0, "", "0123456789")
//
// If more than 2 parameters are passed as preReleaseAndBuild, call cause panic.
// Avoid to call New like:
//   New(1, 0, 0, args...)
// Use following instead:
//   New(1, 0, 0, args[0], args[1])
func New(major, minor, patch uint64, preReleaseAndBuild ...string) Ver {
	preRelease := ""
	build := ""
	switch len(preReleaseAndBuild) {
	case 0: // no-op
	case 2:
		build = preReleaseAndBuild[1]
		fallthrough
	case 1:
		preRelease = preReleaseAndBuild[0]
	default:
		panic("sem.New: len(preReleaseAndBuild) > 2")
	}
	return Ver{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: preRelease,
		Build:      build,
	}
}

// Compare returns 0 if passed value is equal to current one.
// If not, it returns -1 (or 1) if passed value is lower (or greater) than current one.
// Function ComparePreRelease is used as part of Compare.
//
// See also: https://semver.org/#spec-item-11
func (v Ver) Compare(ver Ver) int {
	if ver.Major > v.Major {
		return -1
	} else if ver.Major < v.Major {
		return 1
	}
	if ver.Minor > v.Minor {
		return -1
	} else if ver.Minor < v.Minor {
		return 1
	}
	if ver.Patch > v.Patch {
		return -1
	} else if ver.Patch < v.Patch {
		return 1
	}
	return ComparePreRelease(v.PreRelease, ver.PreRelease)
}

// Latest returns the latest version from passed and current one.
// Method Ver.Compare is used to decide this.
func (v Ver) Latest(ver Ver) Ver {
	if v.Compare(ver) == -1 {
		return ver
	}
	return v
}

// Valid check version validity.
// It checks pre-release and build component of version.
// It can return wrapped ErrInvalidPreRelease or ErrInvalidBuild errors.
func (v Ver) Valid() error {
	if v.PreRelease != "" && !preRelease.MatchString(v.PreRelease) {
		return fmt.Errorf("sem.Ver.Valid: %w", ErrInvalidPreRelease)
	}
	if v.Build != "" && !build.MatchString(v.Build) {
		return fmt.Errorf("sem.Ver.Valid: %w", ErrInvalidBuild)
	}
	return nil
}

// IsZero returns true if version is 0.0.0.
func (v Ver) IsZero() bool {
	return v.Major == 0 && v.Minor == 0 && v.Patch == 0 && v.Build == "" && v.PreRelease == ""
}

// Core returns new version without build and pre-release.
func (v Ver) Core() Ver {
	return Ver{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		PreRelease: "",
		Build:      "",
	}
}

// NextMajor returns new version with incremented major and zero minor and patch.
// Returned version has empty PreRelease and Build components.
// It panics if current major is equal to math.MaxUint64.
func (v Ver) NextMajor() Ver {
	newMajor, overflow := bits.Add64(v.Major, 1, 0)
	if overflow != 0 {
		panic("maximum major version exceeded")
	}
	return Ver{
		Major:      newMajor,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
}

// NextMinor returns new version with same major, incremented minor and zero patch.
// Returned version has empty PreRelease and Build components.
// It panics if current minor is equal to math.MaxUint64.
func (v Ver) NextMinor() Ver {
	newMinor, overflow := bits.Add64(v.Minor, 1, 0)
	if overflow != 0 {
		panic("maximum minor version exceeded")
	}
	return Ver{
		Major:      v.Major,
		Minor:      newMinor,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
}

// NextPatch returns new version with same major, same minor and incremented patch.
// Returned version has empty PreRelease and Build components.
// It panics if current patch is equal to math.MaxUint64.
func (v Ver) NextPatch() Ver {
	newPatch, overflow := bits.Add64(v.Patch, 1, 0)
	if overflow != 0 {
		panic("maximum patch version exceeded")
	}
	return Ver{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      newPatch,
		PreRelease: "",
		Build:      "",
	}
}

// MarshalText converts version to text with Formatter.
func (v Ver) MarshalText() ([]byte, error) {
	b, err := Formatter(nil, v, 0)
	if err != nil {
		return nil, fmt.Errorf("sem.Ver.MarshalText: %w", err)
	}
	return b, nil
}

// UnmarshalText using global Parser function.
func (v *Ver) UnmarshalText(data []byte) error {
	ver, err := Parser(data, 0)
	if err != nil {
		return fmt.Errorf("sem.Ver.UnmarshalText: %w", err)
	}
	*v = ver
	return nil
}

// Format is implementation for fmt.Formatter.
//
//   ┌ Verb ┬ Format ───┬ Example ─────┐
//   │ %s   │ Format(0) │ "1.2.3-b+p"  │
//   │ %t   │ FormatTag │ "v1.2.3-b+p" │
func (v Ver) Format(f fmt.State, verb rune) {
	f.Write(v.format(formatByVerb(verb)))
}

// StringTag formats version as string tag.
// If Formatter returns error, StringTag returns same value as DefaultFormatter.
func (v Ver) StringTag() string {
	return string(v.format(FormatTag))
}

// String formats version for string output.
// If Formatter returns error, String returns same value as DefaultFormatter.
func (v Ver) String() string {
	return string(v.format(0))
}

func (v Ver) format(f Format) []byte {
	b, err := Formatter(nil, v, f)
	if err != nil {
		b, _ = DefaultFormatter(nil, v, f)
	}
	return b
}

func formatByVerb(verb rune) Format {
	switch verb {
	case 't':
		return FormatTag
	default:
		return 0
	}
}
