// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

// ParseVersion parses input as version.
// If input is not valid, error is returned.
func ParseVersion(data []byte) (Ver, error) {
	const funcName = "ParseVersion"
	return unmarshalText(funcName, data, formVersion)
}

// ParseTag parses input as tag.
// If input is not valid, error is returned.
func ParseTag(data []byte) (Ver, error) {
	const funcName = "ParseTag"
	return unmarshalText(funcName, data, formTag)
}

// Parse parses input as version or tag.
// If input is not valid, error is returned.
func Parse(data []byte) (Ver, error) {
	const funcName = "Parse"
	return unmarshalText(funcName, data, formVersion|formTag)
}
