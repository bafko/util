// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefaultUnmarshalText(t *testing.T, expected uint64, input string) {
	t.Helper()
	s, err := DefaultUnmarshalText([]byte(input))
	assert.NoErrorf(t, err, "invalid case for input %q", input)
	assert.Equal(t, Size(expected), s, "invalid case for input %q: expected %d bytes", input, expected)
}

func assertDefaultUnmarshalTextError(t *testing.T, error, input string) {
	t.Helper()
	s, err := DefaultUnmarshalText([]byte(input))
	assert.EqualErrorf(t, err, error, "invalid case for input %q", input)
	assert.Zero(t, s, "invalid case for input %q: expected zero", input)
}

func Test_DefaultUnmarshalText(t *testing.T) {
	MaxTextLength = 4
	DisableUnmarshalTextUnit = false
	assertDefaultUnmarshalTextError(t, `size.DefaultUnmarshalText: input too long (5 > 4)`, "xxxxx")

	MaxTextLength = 0
	assertDefaultUnmarshalTextError(t, `size.DefaultUnmarshalText: parsing "x": unable to parse`, "x")
	assertDefaultUnmarshalTextError(t, `size.DefaultUnmarshalText: parsing "1000000000000000000000000000000": strconv.ParseUint: parsing "1000000000000000000000000000000": value out of range`, "1000000000000000000000000000000")

	DisableUnmarshalTextUnit = true
	assertDefaultUnmarshalTextError(t, `size.DefaultUnmarshalText: parsing "1B": unit disabled`, "1B")

	DisableUnmarshalTextUnit = false
	assertDefaultUnmarshalText(t, 0, "0")
	assertDefaultUnmarshalText(t, 0, "0B")
	assertDefaultUnmarshalText(t, 0, "0 B")
	assertDefaultUnmarshalText(t, 1, "1B")
	assertDefaultUnmarshalText(t, 1, "1 B")
	assertDefaultUnmarshalText(t, 1024, "1 KiB")
	assertDefaultUnmarshalText(t, 1000, "1 kB")
	assertDefaultUnmarshalTextError(t, `size.DefaultUnmarshalText: parsing "1048576 EiB": value 1048576 with unit "EiB" is not suitable for uint64`, "1048576 EiB")
}

func assertDefaultUnmarshalJSON(t *testing.T, expected uint64, input string) {
	t.Helper()
	s, err := DefaultUnmarshalJSON([]byte(input))
	assert.NoErrorf(t, err, "invalid case for input %q", input)
	assert.Equal(t, Size(expected), s, "invalid case for input %q: expected %d bytes", input, expected)
}

func assertDefaultUnmarshalJSONError(t *testing.T, error, input string) {
	t.Helper()
	s, err := DefaultUnmarshalJSON([]byte(input))
	assert.EqualErrorf(t, err, error, "invalid case for input %q", input)
	assert.Zero(t, s, "invalid case for input %q: expected zero", input)
}

func Test_DefaultUnmarshalJSON(t *testing.T) {
	DisableUnmarshalJSONObjectForm = false
	DisableUnmarshalJSONStringForm = false
	UnmarshalText = func(data []byte) (Size, error) {
		assert.Equal(t, []byte("10"), data)
		return 10, nil
	}
	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "x": invalid character 'x' looking for beginning of value`, `x`)
	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "false": unexpected type bool`, `false`)

	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "[]": expected number, string or object`, `[]`)
	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "{}": missing value key`, `{}`)
	assertDefaultUnmarshalJSON(t, 10, `{"value":10,"unit":"B"}`)
	DisableUnmarshalJSONObjectForm = true
	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "{}": object form disabled`, `{}`)

	assertDefaultUnmarshalJSON(t, 10, `10`)

	assertDefaultUnmarshalJSON(t, 10, `"10"`)
	DisableUnmarshalJSONStringForm = true
	assertDefaultUnmarshalJSONError(t, `size.DefaultUnmarshalJSON: parsing "\"10\"": string form disabled`, `"10"`)
}

type spacePermutation []rune

func newSpacePermutation(count int) spacePermutation {
	sp := make(spacePermutation, count)
	for i := range sp {
		sp[i] = '\u0000'
	}
	return sp
}

func (s spacePermutation) Next() bool {
	spaces := map[rune]rune{
		'\u0000': '_',
		'_':      ' ',
		' ':      '\u00A0',
		'\u00A0': '\u0000',
	}
	for i, r := range s {
		s[i] = spaces[r]
		if s[i] != '\u0000' {
			return true
		}
	}
	return false
}

func (s spacePermutation) Replace(input string) string {
	sb := strings.Builder{}
	i := 0
	for _, r := range input {
		if r == '_' {
			if s[i] != '\u0000' {
				sb.WriteRune(s[i])
			}
			i++
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func Test_spacePermutation(t *testing.T) {
	const factor = 4
	sp := newSpacePermutation(factor)
	sequence := "\u0000_ \u00A0"
	expected := []string(nil)
	for _, s := range sequence {
		expected = append(expected, string(s))
	}
	for f := 1; f < factor; f++ {
		original := expected
		expected = nil
		for _, s := range sequence {
			for _, o := range original {
				expected = append(expected, o+string(s))
			}
		}
	}
	s := strings.Repeat(`_`, factor)
	for {
		assert.Equal(t, strings.Replace(expected[0], "\u0000", "", -1), sp.Replace(s))
		expected = expected[1:]
		if !sp.Next() {
			break
		}
	}
	assert.Len(t, expected, 0)
}

func assertPrepareNumber(t *testing.T, expectedNumber, expectedUnit, input string) {
	t.Helper()
	sp := newSpacePermutation(strings.Count(input, "_"))
	for {
		s := sp.Replace(input)
		n, u := prepareNumber(s)
		assert.Equalf(t, expectedNumber, n, `invalid output %q`, s)
		assert.Equalf(t, expectedUnit, u, `invalid output %q`, s)
		if !sp.Next() {
			break
		}
	}
}

func Test_prepareNumber(t *testing.T) {
	assertPrepareNumber(t, `100`, `B`, `100_B`)
	assertPrepareNumber(t, `100`, `B`, ` 1_00B`)
	assertPrepareNumber(t, `100`, `B`, ` 100_B`)
	assertPrepareNumber(t, `100`, `B`, ` 1_00_B`)
	assertPrepareNumber(t, `100`, `B`, `100_B `)
	assertPrepareNumber(t, `100`, `B`, `1_00_B `)
	assertPrepareNumber(t, `100`, `B`, ` 1_00B `)
	assertPrepareNumber(t, `100`, `B`, ` 100_B `)
	assertPrepareNumber(t, `100`, `B`, ` 10_0_B `)
	assertPrepareNumber(t, `100`, `B`, ` 1_00_B `)
	assertPrepareNumber(t, `100`, `B`, ` 1_0_0_B `)
}

var (
	errNoToken = errors.New(`no token`)
)

type decoderMock struct {
	queue []interface{} // json.Token or error
}

func newDecoderMock(tokensOrErrors ...interface{}) *decoderMock {
	return &decoderMock{
		queue: tokensOrErrors,
	}
}

func (d *decoderMock) Token() (json.Token, error) {
	if len(d.queue) == 0 {
		return nil, nil
	}
	t := d.queue[0]
	d.queue = d.queue[1:]
	if err, ok := t.(error); ok {
		return nil, err
	}
	return t, nil
}

func (d *decoderMock) More() bool {
	return len(d.queue) != 0
}

func Test_unmarshalJSONObject(t *testing.T) {
	MaxObjectKeys = 4
	s, err := unmarshalJSONObject(newDecoderMock(errNoToken))
	assert.Zero(t, s)
	assert.EqualError(t, err, `no token`)

	s, err = unmarshalJSONObject(newDecoderMock("key", errNoToken))
	assert.Zero(t, s)
	assert.EqualError(t, err, `no token`)

	s, err = unmarshalJSONObject(newDecoderMock("value", 0))
	assert.Zero(t, s)
	assert.EqualError(t, err, `expected type json.Number instead of int for value`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`)))
	assert.Zero(t, s)
	assert.EqualError(t, err, `missing unit key`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`), errNoToken))
	assert.Zero(t, s)
	assert.EqualError(t, err, `no token`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`), "value"))
	assert.Zero(t, s)
	assert.EqualError(t, err, `duplicated value key`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`), "unit", errNoToken))
	assert.Zero(t, s)
	assert.EqualError(t, err, `no token`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`), "unit", 0))
	assert.Zero(t, s)
	assert.EqualError(t, err, `expected type string instead of int for unit`)

	s, err = unmarshalJSONObject(newDecoderMock("value", json.Number(`0`), "unit", `B`))
	assert.Zero(t, s)
	assert.NoError(t, err)

	s, err = unmarshalJSONObject(newDecoderMock("unit", `B`))
	assert.Zero(t, s)
	assert.EqualError(t, err, `missing value key`)

	s, err = unmarshalJSONObject(newDecoderMock("unit", `B`, "unit"))
	assert.Zero(t, s)
	assert.EqualError(t, err, `duplicated unit key`)

	s, err = unmarshalJSONObject(newDecoderMock("unit", `B`, "value", json.Number(`0`)))
	assert.Zero(t, s)
	assert.NoError(t, err)

	MaxObjectKeys = 2
	s, err = unmarshalJSONObject(newDecoderMock("key", "value", "key", "value", "key", "value"))
	assert.Zero(t, s)
	assert.EqualError(t, err, `object too big (3 > 2)`)
}

func Test_newOrError(t *testing.T) {
	s, err := newOrError(nil, nil)
	assert.Zero(t, s)
	assert.EqualError(t, err, `missing value key`)

	value := uint64(5)
	s, err = newOrError(&value, nil)
	assert.Zero(t, s)
	assert.EqualError(t, err, `missing unit key`)

	unit := Kibibyte
	s, err = newOrError(&value, &unit)
	assert.Equal(t, Size(5*1024), s)
	assert.NoError(t, err)
}

func Test_decodeValue(t *testing.T) {
	s, err := decodeValue(newDecoderMock(errNoToken))
	assert.Nil(t, s)
	assert.EqualError(t, err, `no token`)
	s, err = decodeValue(newDecoderMock(0))
	assert.Nil(t, s)
	assert.EqualError(t, err, `expected type json.Number instead of int for value`)
	s, err = decodeValue(newDecoderMock(json.Number(`x`)))
	assert.Nil(t, s)
	assert.EqualError(t, err, `strconv.ParseUint: parsing "x": invalid syntax`)
	value := uint64(10)
	number := json.Number(`10`)
	s, err = decodeValue(newDecoderMock(number))
	assert.Equal(t, &value, s)
	assert.NoError(t, err)
}

func Test_decodeUnit(t *testing.T) {
	s, err := decodeUnit(newDecoderMock(errNoToken))
	assert.Nil(t, s)
	assert.EqualError(t, err, `no token`)
	s, err = decodeUnit(newDecoderMock(0))
	assert.Nil(t, s)
	assert.EqualError(t, err, `expected type string instead of int for unit`)
	unit := "KiB"
	s, err = decodeUnit(newDecoderMock(unit))
	assert.Equal(t, &unit, s)
	assert.NoError(t, err)
}

func Test_decodeAndSkipNested(t *testing.T) {
	err := decodeAndSkipNested(newDecoderMock(errNoToken))
	assert.EqualError(t, err, `no token`)

	err = decodeAndSkipNested(newDecoderMock(json.Delim('{'), errNoToken))
	assert.EqualError(t, err, `no token`)

	err = decodeAndSkipNested(newDecoderMock(json.Delim('{'), json.Delim('}')))
	assert.NoError(t, err)

	err = decodeAndSkipNested(newDecoderMock(json.Delim('{'), json.Delim('{'), json.Delim('}'), json.Delim('}')))
	assert.NoError(t, err)

	err = decodeAndSkipNested(newDecoderMock(0))
	assert.NoError(t, err)
}
