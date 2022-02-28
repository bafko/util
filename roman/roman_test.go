// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package roman

import (
	"errors"
	"testing"

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
	b, err := Number(15749).MarshalText()
	assert.Equal(t, []byte(`abc`), b)
	assert.NoError(t, err)
	formatError := errors.New("format error")
	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		return nil, formatError
	}
	b, err = Number(15749).MarshalText()
	assert.Nil(t, b)
	assert.Equal(t, formatError, err)
}

func Test_Number_UnmarshalText(t *testing.T) {
	err := errors.New("error")
	Parser = func(input []byte, r Rule) (Number, error) {
		assert.Equal(t, []byte(`abc`), input)
		assert.Equal(t, Rule(0), r)
		return 0, err
	}
	n := Number(123)
	assert.Equal(t, err, n.UnmarshalText([]byte(`abc`)))
	assert.Equal(t, Number(123), n)
	Parser = func(input []byte, r Rule) (Number, error) {
		assert.Equal(t, []byte(`abc`), input)
		assert.Equal(t, Rule(0), r)
		return Number(15749), nil
	}
	assert.NoError(t, n.UnmarshalText([]byte(`abc`)))
	assert.Equal(t, Number(15749), n)
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
	formatError := errors.New("format error")
	Formatter = func(buf []byte, n Number, f Format) ([]byte, error) {
		return nil, formatError
	}
	assert.Equal(t, `MMMMMMMMMMMMMMMDCCXXXXIX`, Number(15749).String())
}
