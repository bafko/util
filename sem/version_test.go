// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package sem

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"go.lstv.dev/util/test"

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

func Test_Ver_Compare(t *testing.T) {
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

func Test_Ver_Latest(t *testing.T) {
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

func Test_Ver_Valid(t *testing.T) {
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

func Test_Ver_IsZero(t *testing.T) {
	v := Ver{}
	assert.True(t, v.IsZero())
	v = Ver{Major: 1}
	assert.False(t, v.IsZero())
	v = Ver{Minor: 1}
	assert.False(t, v.IsZero())
	v = Ver{Patch: 1}
	assert.False(t, v.IsZero())
	v = Ver{Build: "a"}
	assert.False(t, v.IsZero())
	v = Ver{PreRelease: "b"}
	assert.False(t, v.IsZero())
}

func Test_Ver_Core(t *testing.T) {
	v := Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v.Core())
	v.Build = "a"
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v.Core())
	v.Build = "a"
	v.PreRelease = "b"
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v.Core())
	v.Build = ""
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, v.Core())
}

func Test_Ver_NextMajor(t *testing.T) {
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      0,
		Minor:      1,
		Patch:      2,
		PreRelease: "a",
		Build:      "b",
	}.NextMajor())
	assert.Equal(t, Ver{
		Major:      2,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}.NextMajor())
	assert.Panics(t, func() {
		Ver{
			Major:      math.MaxUint64,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}.NextMajor()
	})
}

func Test_Ver_NextMinor(t *testing.T) {
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      0,
		Minor:      0,
		Patch:      2,
		PreRelease: "a",
		Build:      "b",
	}.NextMinor())
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      1,
		Minor:      1,
		Patch:      3,
		PreRelease: "a",
		Build:      "b",
	}.NextMinor())
	assert.Panics(t, func() {
		Ver{
			Major:      0,
			Minor:      math.MaxUint64,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}.NextMinor()
	})
}

func Test_Ver_NextPatch(t *testing.T) {
	assert.Equal(t, Ver{
		Major:      0,
		Minor:      0,
		Patch:      1,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "a",
		Build:      "b",
	}.NextPatch())
	assert.Equal(t, Ver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		PreRelease: "",
		Build:      "",
	}, Ver{
		Major:      1,
		Minor:      2,
		Patch:      2,
		PreRelease: "a",
		Build:      "b",
	}.NextPatch())
	assert.Panics(t, func() {
		Ver{
			Major:      0,
			Minor:      0,
			Patch:      math.MaxUint64,
			PreRelease: "a",
			Build:      "b",
		}.NextPatch()
	})
}

func Test_Ver_MarshalText(t *testing.T) {
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
	test.MarshalText(t, []test.CaseText[Ver]{
		{
			Data: `ab`,
			Value: Ver{
				Major:      2,
				Minor:      1,
				Patch:      3,
				PreRelease: "a",
				Build:      "b",
			},
		},
	})

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
		return nil, errors.New("format error")
	}
	test.MarshalText(t, []test.CaseText[Ver]{
		{
			Error: test.Error("sem.Ver.MarshalText: format error"),
			Value: Ver{
				Major:      2,
				Minor:      1,
				Patch:      3,
				PreRelease: "a",
				Build:      "b",
			},
		},
	})
}

func Test_Ver_UnmarshalText(t *testing.T) {
	Parser = func(input []byte, r Rule) (v Ver, err error) {
		assert.Equal(t, []byte(`ab`), input)
		assert.Equal(t, Rule(0), r)
		return Ver{
			Major:      1,
			Minor:      2,
			Patch:      3,
			PreRelease: "a",
			Build:      "b",
		}, nil
	}
	test.UnmarshalText(t, []test.CaseText[Ver]{
		{
			Data: `ab`,
			Value: Ver{
				Major:      1,
				Minor:      2,
				Patch:      3,
				PreRelease: "a",
				Build:      "b",
			},
		},
	}, nil)

	Parser = func(input []byte, r Rule) (v Ver, err error) {
		assert.Equal(t, []byte(`ab`), input)
		assert.Equal(t, Rule(0), r)
		return Ver{}, errors.New("parse error")
	}
	test.UnmarshalText(t, []test.CaseText[Ver]{
		{
			Error: test.Error("sem.Ver.UnmarshalText: parse error"),
			Data:  `ab`,
		},
	}, nil)
}

func Test_Ver_Format(t *testing.T) {
	Formatter = DefaultFormatter
	assert.Equal(t, "sem.Ver", fmt.Sprintf("%T", New(1, 2, 3, "b", "p")))
	assert.Equal(t, "v1.2.3-b+p", fmt.Sprintf("%t", New(1, 2, 3, "b", "p")))
	assert.Equal(t, "1.2.3-b+p", fmt.Sprintf("%s", New(1, 2, 3, "b", "p")))
}

func Test_Ver_StringTag(t *testing.T) {
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

func Test_Ver_String(t *testing.T) {
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
		return nil, errors.New("format error")
	}
	assert.Equal(t, `1.2.3-a+b`, v.String())
}

func Test_Ver_format(t *testing.T) {
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
	assert.Equal(t, []byte(`x`), v.format(FormatTag))

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
		return nil, errors.New("format error")
	}
	assert.Equal(t, []byte(`v1.2.3-a+b`), v.format(FormatTag))
}

func Test_formatByVerb(t *testing.T) {
	assert.Equal(t, Format(0), formatByVerb(' '))
	assert.Equal(t, Format(0), formatByVerb('s'))
	assert.Equal(t, FormatTag, formatByVerb('t'))
}
