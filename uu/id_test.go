// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package uu

import (
	"errors"
	"fmt"
	"testing"

	"go.lstv.dev/util/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func id(t *testing.T, input string) ID {
	var i ID
	require.NoError(t, i.UnmarshalText([]byte(input)))
	return i
}

func Test_ID_Version(t *testing.T) {
	for i := 0; i <= 15; i++ {
		uuid := fmt.Sprintf("ed7059f3-8044-%xf2a-81aa-b959b33c7777", i)
		assert.Equalf(t, i, id(t, uuid).Version(), "uuid %s with version %d", uuid, i)
	}
}

func Test_ID_Variant(t *testing.T) {
	var expectedVariants = []int{
		0, // 0000
		0, // 0001
		0, // 0010
		0, // 0011
		0, // 0100
		0, // 0101
		0, // 0110
		0, // 0111
		1, // 1000
		1, // 1001
		1, // 1010
		1, // 1011
		2, // 1100
		2, // 1101
		3, // 1110
		3, // 1111
	}
	for i := 0; i <= 15; i++ {
		uuid := fmt.Sprintf("ed7059f3-8044-0f2a-%x1aa-b959b33c7777", i)
		assert.Equalf(t, expectedVariants[i], id(t, uuid).Variant(), "uuid %s with variant %d", uuid, expectedVariants[i])
	}
}

func Test_ID_URN(t *testing.T) {
	assert.Equal(t, `urn:uuid:00000000-0000-0000-0000-000000000000`, ID{}.URN())
	id := ID{
		Higher: 0xed7059f380444f2a,
		Lower:  0x81aab959b33c7777,
	}
	assert.Equal(t, `urn:uuid:ed7059f3-8044-4f2a-81aa-b959b33c7777`, id.URN())
}

func Test_ID_MarshalText(t *testing.T) {
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return []byte(`10-20`), nil
	}
	test.MarshalText(t, []test.CaseText[ID]{
		{
			Data: `10-20`,
			Value: ID{
				Higher: 10,
				Lower:  20,
			},
		},
	})

	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return nil, errors.New("format error")
	}
	test.MarshalText(t, []test.CaseText[ID]{
		{
			Error: test.Error("uu.ID.MarshalText: format error"),
			Value: ID{
				Higher: 10,
				Lower:  20,
			},
		},
	})
}

func Test_ID_UnmarshalText(t *testing.T) {
	text := `10-20`
	value := ID{
		Higher: 10,
		Lower:  20,
	}

	Parser = func(input []byte, r Rule) (ID, error) {
		assert.Equal(t, text, string(input))
		assert.Equal(t, Rule(0), r)
		return value, nil
	}
	test.UnmarshalText(t, []test.CaseText[ID]{
		{
			Data:  text,
			Value: value,
		},
	}, nil)

	Parser = func(input []byte, r Rule) (ID, error) {
		assert.Equal(t, text, string(input))
		assert.Equal(t, Rule(0), r)
		return ID{}, errors.New("parse error")
	}
	test.UnmarshalText(t, []test.CaseText[ID]{
		{
			Error: test.Error("uu.ID.UnmarshalText: parse error"),
			Data:  text,
		},
	}, nil)
}

func Test_ID_Format(t *testing.T) {
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, FormatURN, f)
		return []byte(`urn:uuid:10-20`), nil
	}
	assert.Equal(t, `urn:uuid:10-20`, fmt.Sprintf("%u", ID{
		Higher: 10,
		Lower:  20,
	}))

	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return []byte(`urn:uuid:10-20`), nil
	}
	assert.Equal(t, `urn:uuid:10-20`, fmt.Sprintf("%s", ID{
		Higher: 10,
		Lower:  20,
	}))

	formatError := errors.New("format error")
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, FormatURN, f)
		return nil, formatError
	}
	assert.Equal(t, `urn:uuid:00000000-0000-000a-0000-000000000014`, fmt.Sprintf("%u", ID{
		Higher: 10,
		Lower:  20,
	}))
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return nil, formatError
	}
	assert.Equal(t, `00000000-0000-000a-0000-000000000014`, fmt.Sprintf("%s", ID{
		Higher: 10,
		Lower:  20,
	}))
}

func Test_ID_String(t *testing.T) {
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return []byte(`urn:uuid:10-20`), nil
	}
	assert.Equal(t, `urn:uuid:10-20`, ID{
		Higher: 10,
		Lower:  20,
	}.String())

	formatError := errors.New("format error")
	Formatter = func(buf []byte, id ID, f Format) ([]byte, error) {
		assert.Equal(t, ID{
			Higher: 10,
			Lower:  20,
		}, id)
		assert.Equal(t, Format(0), f)
		return nil, formatError
	}
	assert.Equal(t, `00000000-0000-000a-0000-000000000014`, ID{
		Higher: 10,
		Lower:  20,
	}.String())
}

func Test_ID(t *testing.T) {
	Formatter = DefaultFormatter
	Parser = DefaultParser[[]byte]

	var i ID
	require.Error(t, i.UnmarshalText(nil))
	require.Error(t, i.UnmarshalText([]byte("ed7059f3x8044x4f2ax81aaxb959b33c7777")))
	require.Error(t, i.UnmarshalText([]byte("ed7059fx-8044-4f2a-81aa-b959b33c7777")))
	v := []byte("ed7059f3-8044-4f2a-81aa-b959b33c7777")
	require.NoError(t, i.UnmarshalText(v))
	b, err := i.MarshalText()
	require.NoError(t, err)
	assert.Equal(t, v, b)
}
