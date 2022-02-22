// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"regexp"
)

const (
	tagPrefix = 'v'

	digitPattern           = `[0-9]`
	digitsPattern          = digitPattern + `+`
	nonDigitPattern        = `[A-Za-z\-]`
	identPattern           = `[0-9A-Za-z\-]`
	numIdentPattern        = `0|[1-9][0-9]*`
	alphanumIdentPattern   = `(?:` + identPattern + `*` + nonDigitPattern + identPattern + `*)`
	versionCorePattern     = `(` + numIdentPattern + `)\.(` + numIdentPattern + `)\.(` + numIdentPattern + `)`
	preReleasePattern      = preReleaseIdentPattern + `(?:\.` + preReleaseIdentPattern + `)*`
	preReleaseIdentPattern = `(?:` + alphanumIdentPattern + `|(?:` + numIdentPattern + `))`
	buildPattern           = buildIdentPattern + `(?:\.` + buildIdentPattern + `)*`
	buildIdentPattern      = `(?:` + alphanumIdentPattern + `|` + digitsPattern + `)`
	semverPattern          = `^` + versionCorePattern + `(?:\-(` + preReleasePattern + `))?(?:\+(` + buildPattern + `))?$`
)

var (
	pattern       = regexp.MustCompile(semverPattern)
	digitsOrEmpty = regexp.MustCompile(`^` + digitPattern + `*$`)
	preRelease    = regexp.MustCompile(`^` + preReleasePattern + `$`)
	build         = regexp.MustCompile(`^` + buildPattern + `$`)
)
