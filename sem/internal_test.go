// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var boolToValid = map[bool]string{
	false: "invalid",
	true:  "valid",
}

func assertRegexp(t *testing.T, r *regexp.Regexp, valid bool, v string) {
	assert.Equalf(t, valid, r.MatchString(v), "expected %q as %s %q", v, boolToValid[valid], r.String())
}

func Test_versionCorePattern(t *testing.T) {
	r := regexp.MustCompile(`^` + versionCorePattern + `$`)
	for major := 0; major < 12; major++ {
		for minor := 0; minor < 12; minor++ {
			for patch := 0; patch < 12; patch++ {
				v := fmt.Sprintf("%d.%d.%d", major, minor, patch)
				assertRegexp(t, r, true, v)
			}
		}
	}
	invalidVersions := []string{
		"1",
		"1.0",
		"1.0.0.0",
		"01.0.0",
		"1.1.1a",
	}
	for _, v := range invalidVersions {
		assertRegexp(t, r, false, v)
	}
}

func Test_alphanumIdentPattern(t *testing.T) {
	r := regexp.MustCompile(`^` + alphanumIdentPattern + `$`)
	valid := []string{
		"a",
		"A",
		"-",

		"1a",
		"1A",
		"1-",
		"a1",
		"A1",
		"-1",
		"1a1",
		"1A1",
		"1-1",

		"aa",
		"aA",
		"a-",
		"aa",
		"Aa",
		"-a",
		"aaa",
		"aAa",
		"a-a",

		"Aa",
		"AA",
		"A-",
		"aA",
		"AA",
		"-A",
		"AaA",
		"AAA",
		"A-A",

		"-a",
		"-A",
		"--",
		"a-",
		"A-",
		"--",
		"-a-",
		"-A-",
		"---",
	}
	invalid := []string{
		"*",
		"1",
		"11",
	}
	for _, v := range valid {
		assertRegexp(t, r, true, v)
	}
	for _, v := range invalid {
		assertRegexp(t, r, false, v)
	}
}
