// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package size

import (
	"encoding/json"
	"errors"
	"html/template"
	"math"
	"testing"

	"go.lstv.dev/util/constraint"
	"go.lstv.dev/util/test"

	"github.com/stretchr/testify/assert"
)

func assertNew[N constraint.Numbers](t *testing.T, expected Size, value N, unit string) {
	t.Helper()
	actual, err := New[N](value, unit)
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func assertNewFail[N constraint.Numbers](t *testing.T, error string, value N, unit string) {
	t.Helper()
	actual, err := New[N](value, unit)
	assert.Zero(t, actual)
	assert.EqualError(t, err, error)
}

func Test_New(t *testing.T) {
	assertNew(t, 0, 0, "")

	assertNewFail(t, `invalid unit "h"`, 0, "h")
	assertNew(t, 1, 1, "")
	assertNewFail(t, `invalid unit "h"`, 1, "h")
	assertNewFail(t, `invalid unit "YB"`, 1, Yottabyte)
	assertNewFail(t, `invalid unit "YiB"`, 1, Yobibyte)
	assertNewFail(t, `invalid unit "ZB"`, 1, Zettabyte)
	assertNewFail(t, `invalid unit "ZiB"`, 1, Zebibyte)
	assertNewFail(t, `value 18446744073709551615 with unit "EiB" is not suitable for uint64`, uint64(math.MaxUint64), Exbibyte)
	assertNew(t, 1024*1024, 1, Mebibyte)

	assertNew(t, 10000000, int(10000000), "")
	assertNew(t, 0, float32(0), "")
	assertNew(t, 0, float64(0), "")
	assertNew(t, 15, float64(15), "")
	assertNewFail(t, `value -1 without unit is not suitable for uint64`, int(-1), "")
	assertNewFail(t, `value 1.8446744073709552e+21 without unit is not suitable for uint64`, float64(math.MaxUint64)*100, "")
	assertNewFail(t, `value 0.3 without unit is not suitable for uint64`, float64(0.3), "")
}

func Test_Size_Shorten(t *testing.T) {
	cases := []struct {
		size  Size
		value uint64
		unit  string
	}{
		{
			size:  0,
			value: 0,
			unit:  Byte,
		},
		{
			size:  1023,
			value: 1023,
			unit:  Byte,
		},
		{
			size:  1024,
			value: 1,
			unit:  Kibibyte,
		},
		{
			size:  1025,
			value: 1025,
			unit:  Byte,
		},
		{
			size:  1024 * 1024,
			value: 1,
			unit:  Mebibyte,
		},
		{
			size:  1024 * 1024 * 1024,
			value: 1,
			unit:  Gibibyte,
		},
		{
			size:  1024 * 1024 * 1024 * 1024,
			value: 1,
			unit:  Tebibyte,
		},
		{
			size:  1024 * 1024 * 1024 * 1024 * 1024,
			value: 1,
			unit:  Pebibyte,
		},
		{
			size:  1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			value: 1,
			unit:  Exbibyte,
		},
		{
			size:  math.MaxUint64,
			value: math.MaxUint64,
			unit:  Byte,
		},
	}
	for i, c := range cases {
		v, u := c.size.Shorten()
		assert.Equalf(t, c.value, v, "invalid case %d", i)
		assert.Equal(t, c.unit, u, "invalid case %d", i)
	}
}

func Test_Size_BytesJSONNumber(t *testing.T) {
	assert.Equal(t, json.Number(`10`), Size(10).BytesJSONNumber())
}

func Test_Size_BytesString(t *testing.T) {
	assert.Equal(t, `10`, Size(10).BytesString())
}

func Test_Size_MarshalText(t *testing.T) {
	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return []byte(`+10`), nil
	}
	test.MarshalText(t, []test.CaseText[Size]{
		{
			Data:  `+10`,
			Value: Size(10),
		},
	})

	DisableMarshalTextUnit = true
	test.MarshalText(t, []test.CaseText[Size]{
		{
			Data:  `10`,
			Value: Size(10),
		},
	})

	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return nil, errors.New("format error")
	}
	test.MarshalText(t, []test.CaseText[Size]{
		{
			Error: test.Error("format error"),
			Value: Size(10),
		},
	})
}

func Test_Size_UnmarshalText(t *testing.T) {
	text := `20KiB`
	value := Size(20 * 1024)

	Parser = func(input []byte, r Rule) (Size, error) {
		assert.Equal(t, text, string(input))
		assert.Equal(t, DefaultRule&ruleUnmarshalTextMask, r)
		return value, nil
	}
	test.UnmarshalText(t, []test.CaseText[Size]{
		{
			Data:  text,
			Value: value,
		},
	}, nil)

	Parser = func(input []byte, r Rule) (Size, error) {
		assert.Equal(t, text, string(input))
		assert.Equal(t, DefaultRule&ruleUnmarshalTextMask, r)
		return 0, errors.New("parse error")
	}
	test.UnmarshalText(t, []test.CaseText[Size]{
		{
			Error: test.Error("parse error"),
			Data:  text,
		},
	}, nil)
}

func Test_Size_MarshalJSON(t *testing.T) {
	DisableMarshalJSONObjectForm = false
	DisableMarshalJSONStringForm = false
	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return []byte(`+10`), nil
	}
	test.MarshalJSON(t, []test.CaseJSON[Size]{
		{
			Data:  `{"value":10,"unit":"B"}`,
			Value: Size(10),
		},
	})

	DisableMarshalJSONObjectForm = true
	test.MarshalJSON(t, []test.CaseJSON[Size]{
		{
			Data:  `"+10"`,
			Value: Size(10),
		},
	})

	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return nil, errors.New("format error")
	}
	test.MarshalJSON(t, []test.CaseJSON[Size]{
		{
			Error: test.Error("format error"),
			Value: Size(10),
		},
	})

	DisableMarshalJSONStringForm = true
	test.MarshalJSON(t, []test.CaseJSON[Size]{
		{
			Data:  `10`,
			Value: Size(10),
		},
	})
}

func Test_Size_UnmarshalJSON(t *testing.T) {
	json := `"20KiB"`
	value := Size(20 * 1024)

	Parser = func(input []byte, r Rule) (Size, error) {
		assert.Equal(t, json, string(input))
		assert.Equal(t, DefaultRule, r)
		return value, nil
	}
	test.UnmarshalJSON(t, []test.CaseJSON[Size]{
		{
			Data:  json,
			Value: value,
		},
	}, nil)

	Parser = func(input []byte, r Rule) (Size, error) {
		assert.Equal(t, json, string(input))
		assert.Equal(t, DefaultRule, r)
		return 0, errors.New("parse error")
	}
	test.UnmarshalJSON(t, []test.CaseJSON[Size]{
		{
			Error: test.Error("parse error"),
			Data:  json,
		},
	}, nil)
}

func Test_Size_PrettyHTML(t *testing.T) {
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, FormatPretty|FormatHTML, f)
		return []byte(`+10`), nil
	}
	assert.Equal(t, template.HTML(`+10`), Size(10).PrettyHTML())

	formatError := errors.New("format error")
	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, FormatPretty|FormatHTML, f)
		return nil, formatError
	}
	assert.Panics(t, func() {
		Size(10).PrettyHTML()
	})
}

func Test_Size_PrettyString(t *testing.T) {
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, FormatPretty, f)
		return []byte(`+10`), nil
	}
	assert.Equal(t, `+10`, Size(10).PrettyString())

	formatError := errors.New("format error")
	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, FormatPretty, f)
		return nil, formatError
	}
	assert.Panics(t, func() {
		Size(10).PrettyString()
	})
}

func Test_Size_String(t *testing.T) {
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return []byte(`+10`), nil
	}
	assert.Equal(t, `+10`, Size(10).String())

	formatError := errors.New("format error")
	DisableMarshalTextUnit = false
	Formatter = func(buf []byte, s Size, f Format) ([]byte, error) {
		assert.Equal(t, Size(10), s)
		assert.Equal(t, Format(0), f)
		return nil, formatError
	}
	assert.Equal(t, `10`, Size(10).String())
}

func Test_Size_marshalJSONObject(t *testing.T) {
	assert.Equal(t, []byte(`{"value":1,"unit":"KiB"}`), Size(1024).marshalJSONObject())
}

func assertBytes[N constraint.Numbers](t *testing.T, expected N, size Size) {
	t.Helper()
	actual, ok := Bytes[N](size)
	assert.Equal(t, expected, actual)
	assert.True(t, ok)
}

func assertBytesFail[N constraint.Numbers](t *testing.T, size Size) {
	t.Helper()
	actual, ok := Bytes[N](size)
	assert.Zero(t, actual)
	assert.False(t, ok)
}

func Test_Bytes(t *testing.T) {
	assertBytes[int](t, 0, 0)
	assertBytes[int8](t, 0, 0)
	assertBytes[int16](t, 0, 0)
	assertBytes[int32](t, 0, 0)
	assertBytes[int64](t, 0, 0)
	assertBytes[uint](t, 0, 0)
	assertBytes[uint8](t, 0, 0)
	assertBytes[uint16](t, 0, 0)
	assertBytes[uint32](t, 0, 0)
	assertBytes[uint64](t, 0, 0)
	assertBytes[float32](t, 0, 0)
	assertBytes[float64](t, 0, 0)
	assertBytesFail[int8](t, math.MaxInt8+1)
	assertBytesFail[int16](t, math.MaxInt16+1)
	assertBytesFail[int32](t, math.MaxInt32+1)
	assertBytesFail[int64](t, math.MaxInt64+1)
	assertBytesFail[uint8](t, math.MaxUint8+1)
	assertBytesFail[uint16](t, math.MaxUint16+1)
	assertBytesFail[uint32](t, math.MaxUint32+1)
	assertBytesFail[float32](t, math.MaxUint64)
	assertBytesFail[float64](t, math.MaxUint64)
}
