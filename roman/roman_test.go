// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"errors"
	"fmt"
	"testing"

	"go.lstv.dev/util/test"

	"github.com/stretchr/testify/assert"
)

func Test_Number_MarshalText(t *testing.T) {
	DefaultFormat = FormatLong40
	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Number(15749), n)
		assert.Equal(t, DefaultFormat, f)
		return []byte(`abc`), nil
	}
	test.MarshalText(t, []test.CaseText[Number]{
		{
			Data:  `abc`,
			Value: Number(15749),
		},
	})

	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		return nil, errors.New("format error")
	}
	test.MarshalText(t, []test.CaseText[Number]{
		{
			Error: test.Error("format error"),
			Value: Number(15749),
		},
	})
}

func Test_Number_UnmarshalText(t *testing.T) {
	Parser = func(input []byte, r Rule) (Number, error) {
		assert.Equal(t, []byte(`abc`), input)
		assert.Equal(t, Rule(0), r)
		return 0, errors.New("error")
	}
	test.UnmarshalText(t, []test.CaseText[Number]{
		{
			Error: test.Error("error"),
			Data:  `abc`,
		},
	}, nil)

	Parser = func(input []byte, r Rule) (Number, error) {
		assert.Equal(t, []byte(`abc`), input)
		assert.Equal(t, Rule(0), r)
		return Number(15749), nil
	}
	test.UnmarshalText(t, []test.CaseText[Number]{
		{
			Data:  `abc`,
			Value: Number(15749),
		},
	}, nil)
}

func Test_Number_Format(t *testing.T) {
	Formatter = DefaultFormatter
	n := Number(4)
	assert.Equal(t, `roman.Number`, fmt.Sprintf("%T", n))
	assert.Equal(t, `IIII`, fmt.Sprintf("%L", n))
	assert.Equal(t, `iiii`, fmt.Sprintf("%l", n))
	assert.Equal(t, `IV`, fmt.Sprintf("%R", n))
	assert.Equal(t, `iv`, fmt.Sprintf("%r", n))
	assert.Equal(t, `IV`, fmt.Sprintf("%s", n))
}

func Test_Number_String(t *testing.T) {
	DefaultFormat = FormatLong40
	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Number(15749), n)
		assert.Equal(t, DefaultFormat, f)
		return []byte(`abc`), nil
	}
	assert.Equal(t, `abc`, Number(15749).String())

	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		return nil, errors.New("format error")
	}
	assert.Equal(t, `MMMMMMMMMMMMMMMDCCXXXXIX`, Number(15749).String())
}

func Test_Number_format(t *testing.T) {
	DefaultFormat = 0
	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Number(15749), n)
		assert.Equal(t, FormatLong40, f)
		return []byte(`abc`), nil
	}
	assert.Equal(t, []byte(`abc`), Number(15749).format(FormatLong40))

	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		return nil, errors.New("format error")
	}
	assert.Equal(t, `MMMMMMMMMMMMMMMDCCXLIX`, Number(15749).String())
}

func Test_formatByVerb(t *testing.T) {
	DefaultFormat = FormatLong40
	assert.Equal(t, DefaultFormat, formatByVerb(' '))
	assert.Equal(t, DefaultFormat, formatByVerb('s'))
	assert.Equal(t, FormatLowerCase, formatByVerb('r'))
	assert.Equal(t, Format(0), formatByVerb('R'))
	assert.Equal(t, FormatLowerCase|FormatLong, formatByVerb('l'))
	assert.Equal(t, FormatLong, formatByVerb('L'))
}
