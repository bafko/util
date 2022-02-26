// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	assert.Equal(t, Ver{}, New(0, 0, 0))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, New(1, 2, 3))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "",
	}, New(1, 2, 3, "x"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "x",
		Build:      "y",
	}, New(1, 2, 3, "x", "y"))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "y",
	}, New(1, 2, 3, "", "y"))
	assert.Panics(t, func() {
		New(1, 2, 3, "x", "y", "z")
	})
}

func Test_Version_Compare(t *testing.T) {
	a := Ver{}
	b := Ver{}
	assert.Equal(t, 0, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      1,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, 1, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      1,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, -1, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, 1, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, -1, a.Compare(b))
	a = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, 1, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, -1, a.Compare(b))
	a = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      5,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, 1, a.Compare(b))
	a = Ver{
		Major:      0,
		Minor:      5,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, -1, a.Compare(b))
	a = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "alfa.1",
		Build:      "",
	}
	b = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, -1, a.Compare(b))
	assert.Equal(t, 1, b.Compare(a))
	b.PreRelease = "alfa.1"
	assert.Equal(t, 0, a.Compare(b))
	assert.Equal(t, 0, b.Compare(a))
	b.PreRelease = "alfa.2"
	assert.Equal(t, -1, a.Compare(b))
	assert.Equal(t, 1, b.Compare(a))
}

func Test_Version_Latest(t *testing.T) {
	a := Ver{}
	b := Ver{}
	assert.Equal(t, a, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      1,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, a, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      1,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, b, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, a, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, b, a.Latest(b))
	a = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, a, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      0,
		Patch:      5,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, b, a.Latest(b))
	a = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      0,
		Minor:      5,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, a, a.Latest(b))
	a = Ver{
		Major:      0,
		Minor:      5,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	b = Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, b, a.Latest(b))
}

func Test_Version_Valid(t *testing.T) {
	v := Ver{}
	assert.NoError(t, v.Valid())
	v.PreRelease = "abcd"
	assert.NoError(t, v.Valid())
	v.Build = "123"
	assert.NoError(t, v.Valid())
	v.PreRelease = ""
	assert.NoError(t, v.Valid())
	v.PreRelease = "*"
	assert.Error(t, v.Valid())
	v.Build = ""
	assert.Error(t, v.Valid())
	v.Build = "*"
	assert.Error(t, v.Valid())
	v.PreRelease = "0"
	assert.Error(t, v.Valid())
	v.PreRelease = "abcd"
	assert.Error(t, v.Valid())
}

func Test_Version_MarshalText(t *testing.T) {
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      2,
			Minor:      1,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, Format(0), f)
		return []byte(`ab`), nil
	}
	b, err := Ver{
		Major:      2,
		Minor:      1,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}.MarshalText()
	assert.Equal(t, []byte(`ab`), b)
	assert.NoError(t, err)
	formatError := errors.New("format error")
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      2,
			Minor:      1,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, Format(0), f)
		return nil, formatError
	}
	b, err = Ver{
		Major:      2,
		Minor:      1,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}.MarshalText()
	assert.Nil(t, b)
	assert.Equal(t, formatError, err)
}

func Test_Version_UnmarshalText(t *testing.T) {
	Parser = func(data []byte, r Rule) (v Ver, err error) {
		assert.Equal(t, []byte(`ab`), data)
		assert.Equal(t, Rule(0), r)
		return Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, nil
	}
	v := Ver{}
	assert.NoError(t, v.UnmarshalText([]byte(`ab`)))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}, v)
	parseError := errors.New("parse error")
	Parser = func(data []byte, r Rule) (v Ver, err error) {
		assert.Equal(t, []byte(`ab`), data)
		assert.Equal(t, Rule(0), r)
		return Ver{}, parseError
	}
	assert.Equal(t, parseError, v.UnmarshalText([]byte(`ab`)))
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}, v)
}

func Test_Version_StringTag(t *testing.T) {
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, FormatTag, f)
		return []byte(`x`), nil
	}
	v := Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}
	assert.Equal(t, `x`, v.StringTag())
	formatError := errors.New("format error")
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, FormatTag, f)
		return nil, formatError
	}
	assert.Equal(t, `v1.2.3-a+b`, v.StringTag())
}

func Test_Version_String(t *testing.T) {
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, Format(0), f)
		return []byte(`x`), nil
	}
	v := Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}
	assert.Equal(t, `x`, v.String())
	formatError := errors.New("format error")
	Formatter = func(buf []byte, v Ver, f Format) ([]byte, error) {
		assert.Nil(t, buf)
		assert.Equal(t, Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, v)
		assert.Equal(t, Format(0), f)
		return nil, formatError
	}
	assert.Equal(t, `1.2.3-a+b`, v.String())
}
