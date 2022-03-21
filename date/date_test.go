// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"encoding/xml"
	"errors"
	"fmt"
	"testing"
	"time"

	"go.lstv.dev/util/test"

	"github.com/stretchr/testify/assert"
)

func assertDate(t *testing.T, expectedYear int, expectedMonth Month, expectedDay int, date Date) {
	t.Helper()
	year, month, day := date.Date()
	assert.Equal(t, expectedYear, year)
	assert.Equal(t, expectedMonth, month)
	assert.Equal(t, expectedDay, day)
}

func ExampleDate() {
	d := New(2022, January, 30)
	fmt.Println(d.String())
	// Output: 2022-01-30
}

func Test_Date(t *testing.T) {
	assertDate(t, 1, January, 1, Date{})
	assertDate(t, 2002, August, 7, New(2002, August, 7))
	assertDate(t, 2002, August, 7, FromTime(time.Date(2002, August, 7, 15, 12, 55, 7, time.UTC)))
}

func Test_Today(t *testing.T) {
	date := Today()
	year, month, day := time.Now().Date()
	now := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	assert.True(t, now.Sub(date.Time()) >= 0)
}

func Test_Date_IsZero(t *testing.T) {
	assert.False(t, New(2002, August, 7).IsZero())
	assert.True(t, Date{}.IsZero())
	assert.True(t, FromTime(time.Time{}).IsZero())
}

func Test_Date_Equal(t *testing.T) {
	assert.False(t, New(2002, August, 7).Equal(FromTime(time.Time{})))
	assert.True(t, Date{}.Equal(FromTime(time.Time{})))
}

func Test_Date_Date(t *testing.T) {
	date := New(2002, August, 7)
	year, month, day := date.Date()
	assert.Equal(t, 2002, year)
	assert.Equal(t, August, month)
	assert.Equal(t, 7, day)
	assert.Equal(t, 2002, date.Year())
	assert.Equal(t, August, date.Month())
	assert.Equal(t, 7, date.Day())
}

func Test_Date_Time(t *testing.T) {
	expectedTime := time.Date(2002, August, 7, 0, 0, 0, 0, time.UTC)
	timeWithLocation := time.Date(2002, August, 7, 14, 12, 55, 7, time.FixedZone("+1", 60*60))
	assert.Equal(t, expectedTime, FromTime(timeWithLocation).Time())
}

func Test_Date_FromTime(t *testing.T) {
	timeWithLocation := time.Date(2002, August, 7, 14, 12, 55, 7, time.FixedZone("+1", 60*60))
	date := Date{}
	date.FromTime(timeWithLocation)
	assertDate(t, 2002, August, 7, date)
}

func Test_Date_Add(t *testing.T) {
	assertDate(t, 2004, October, 8, New(2002, August, 7).Add(2, 2, 1))
	assertDate(t, 2004, April, 7, New(2002, August, 7).Add(0, 20, 0))
	assertDate(t, 2000, December, 7, New(2002, August, 7).Add(0, -20, 0))
}

func Test_Date_AddDuration(t *testing.T) {
	assertDate(t, 2002, August, 7, New(2002, August, 7).AddDuration(time.Hour))
	assertDate(t, 2002, August, 8, New(2002, August, 7).AddDuration(time.Hour*24))
	assertDate(t, 2002, August, 6, New(2002, August, 7).AddDuration(-time.Hour*24))
}

func Test_Date_After(t *testing.T) {
	assert.False(t, New(2002, August, 7).After(New(2004, August, 7)))
	assert.False(t, New(2002, August, 7).After(New(2002, October, 7)))
	assert.False(t, New(2002, August, 7).After(New(2002, August, 9)))
	assert.False(t, New(2002, August, 7).After(New(2002, August, 7)))
	assert.True(t, New(2002, August, 7).After(New(2002, August, 5)))
	assert.True(t, New(2002, August, 7).After(New(2000, August, 7)))
	assert.True(t, New(2002, August, 7).After(New(2002, June, 7)))
}

func Test_Date_Before(t *testing.T) {
	assert.False(t, New(2002, August, 7).Before(New(2002, August, 7)))
	assert.False(t, New(2002, August, 7).Before(New(2000, August, 7)))
	assert.False(t, New(2002, August, 7).Before(New(2002, June, 7)))
	assert.False(t, New(2002, August, 7).Before(New(2002, August, 5)))
	assert.True(t, New(2002, August, 7).Before(New(2004, August, 7)))
	assert.True(t, New(2002, August, 7).Before(New(2002, October, 7)))
	assert.True(t, New(2002, August, 7).Before(New(2002, August, 9)))
}

func Test_Date_Sub(t *testing.T) {
	assert.Equal(t, time.Duration(0), New(2002, August, 7).Sub(New(2002, August, 7)))
	assert.Equal(t, time.Hour*24, New(2002, August, 7).Sub(New(2002, August, 6)))
	assert.Equal(t, -time.Hour*24, New(2002, August, 7).Sub(New(2002, August, 8)))
}

func Test_Date_DaysBetween(t *testing.T) {
	assert.Equal(t, 0, New(2002, August, 7).DaysBetween(New(2002, August, 7)))
	assert.Equal(t, 1, New(2002, August, 7).DaysBetween(New(2002, August, 6)))
	assert.Equal(t, -1, New(2002, August, 7).DaysBetween(New(2002, August, 8)))
	// 2012-06-30 has overlap second
	assert.Equal(t, -366, New(2012, January, 1).DaysBetween(New(2013, January, 1)))
}

func Test_Date_MarshalBinary(t *testing.T) {
	test.MarshalBinary(t, []test.CaseBinary[Date]{
		{ // 0
			Data:  []byte{1, 0, 0, 0, 1, 1, 1},
			Value: Date{},
		},
		{ // 1
			Data:  []byte{1, 0, 0, 0x7, 0xd2, 8, 7},
			Value: New(2002, August, 7),
		},
	})
}

func Test_Date_UnmarshalBinary(t *testing.T) {
	test.UnmarshalBinary(t, []test.CaseBinary[Date]{
		{ // 0
			Data:  []byte{1, 0, 0, 0, 1, 1, 1},
			Value: Date{},
		},
		{ // 1
			Data:  []byte{1, 0, 0, 0x7, 0xd2, 8, 7},
			Value: New(2002, August, 7),
		},
		{ // 2
			Error: test.Error("Date.UnmarshalBinary: invalid length: empty data"),
			Data:  []byte(nil),
		},
		{ // 3
			Error: test.Error("Date.UnmarshalBinary: invalid length: expected 7 instead of 1"),
			Data:  []byte{1},
		},
		{ // 4
			Error: test.Error("Date.UnmarshalBinary: unsupported version: expected 1 instead of 50"),
			Data:  []byte(`2002`),
		},
		{ // 5
			Error: test.Error("Date.UnmarshalBinary: unsupported version: expected 1 instead of 34"),
			Data:  []byte(`"2002-08-07"`),
		},
	}, nil)
}

func Test_Date_MarshalText(t *testing.T) {
	test.MarshalText(t, []test.CaseText[Date]{
		{ // 0
			Data:  `0001-01-01`,
			Value: Date{},
		},
		{ // 1
			Data:  `2002-08-07`,
			Value: New(2002, August, 7),
		},
	})
}

func Test_Date_UnmarshalText(t *testing.T) {
	test.UnmarshalText(t, []test.CaseText[Date]{
		{ // 0
			Data:  `0001-01-01`,
			Value: Date{},
		},
		{ // 1
			Data:  `2002-08-07`,
			Value: New(2002, August, 7),
		},
		{ // 2
			Error: test.Error("date.DefaultParser: \"2002\": invalid date"),
			Data:  `2002`,
		},
		{ // 3
			Error: test.Error("date.DefaultParser: input too long: 12 > 10"),
			Data:  `"2002-08-07"`,
		},
	}, nil)
}

func Test_Date_xml(t *testing.T) {
	b := []byte(`<test a="2020-08-05"><b>2020-08-07</b></test>`)
	test := struct {
		A Date `xml:"a,attr"`
		B Date `xml:"b"`
	}{}
	assert.NoError(t, xml.Unmarshal(b, &test))
	assertDate(t, 2020, August, 5, test.A)
	assertDate(t, 2020, August, 7, test.B)
}

func Test_Date_Scan(t *testing.T) {
	date := Date{}
	assert.NoError(t, date.Scan(time.Date(2002, August, 7, 14, 12, 55, 7, time.FixedZone("+1", 60*60))))
	assertDate(t, 2002, August, 7, date)
	assert.Error(t, date.Scan(nil))
}

func Test_Date_Value(t *testing.T) {
	date := New(2002, August, 7)
	v, err := date.Value()
	assert.Equal(t, time.Date(2002, August, 7, 0, 0, 0, 0, time.UTC), v)
	assert.NoError(t, err)
}

func Test_Date_String(t *testing.T) {
	err := errors.New("error")
	Formatter = func(buf []byte, d Date, f Format) ([]byte, error) {
		return nil, err
	}
	assert.Equal(t, `0001-01-01`, Date{}.String())
	assert.Equal(t, `2002-08-07`, New(2002, August, 7).String())
	Formatter = DefaultFormatter
	assert.Equal(t, `0001-01-01`, Date{}.String())
	assert.Equal(t, `2002-08-07`, New(2002, August, 7).String())
}
