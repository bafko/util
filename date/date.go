// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package date provides type Date to keep, marshal and unmarshal date values.
// Date consists of year, month and day.
// Package is compatible with default time package.
package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	// TimeFormatExtended is date format for time package.
	TimeFormatExtended = `2006-01-02`

	// TimeFormatBasic is date format for time package.
	TimeFormatBasic = `20060102`

	version = 1
)

// Date is representation of date.
// Zero date is valid, representing value 0001-01-01.
type Date struct {
	year  int32
	month uint8
	day   uint8
}

// Today returns actual date.
func Today() Date {
	return FromTime(time.Now())
}

// New creates date with specific year, month and day.
func New(year int, month Month, day int) Date {
	return FromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// FromTime creates date from time.Time value.
func FromTime(t time.Time) Date {
	d := Date{}
	d.FromTime(t)
	return d
}

// IsZero returns true if date is 0001-01-01.
func (d Date) IsZero() bool {
	return d.year == 0 && d.month == 0 && d.day == 0
}

// Equal returns true if passed date is the same value.
func (d Date) Equal(e Date) bool {
	return d.year == e.year && d.month == e.month && d.day == e.day
}

// Date returns year, month and day values.
func (d Date) Date() (year int, month Month, day int) {
	return int(d.year + 1), Month(d.month + 1), int(d.day + 1)
}

// Year returns date year.
func (d Date) Year() int {
	return int(d.year + 1)
}

// Month returns date month (from January to December).
func (d Date) Month() Month {
	return Month(d.month + 1)
}

// Day returns date day (1-31).
func (d Date) Day() int {
	return int(d.day + 1)
}

// Time returns time.Time object based on date value.
// Time is midnight (0:00:00.0) and zone is time.UTC.
func (d Date) Time() time.Time {
	return time.Date(int(d.year+1), time.Month(d.month+1), int(d.day+1), 0, 0, 0, 0, time.UTC)
}

// FromTime sets date year, month and day from passed time.Time value.
func (d *Date) FromTime(t time.Time) {
	if t.IsZero() {
		d.year = 0
		d.month = 0
		d.day = 0
		return
	}
	year, month, day := t.Date()
	d.year = int32(year - 1)
	d.month = uint8(month) - 1
	d.day = uint8(day - 1)
}

// Add passed values to date.
func (d Date) Add(years int, months int, days int) Date {
	return FromTime(d.Time().AddDate(years, months, days))
}

// AddDuration add passed duration to date.
func (d Date) AddDuration(duration time.Duration) Date {
	return FromTime(d.Time().Add(duration))
}

// After returns true if passed date is after current one.
// Otherwise, and also if dates are equal, returns false.
func (d Date) After(e Date) bool {
	if d.year > e.year {
		return true
	}
	if d.year < e.year {
		return false
	}
	if d.month > e.month {
		return true
	}
	if d.month < e.month {
		return false
	}
	if d.day > e.day {
		return true
	}
	return false
}

// Before return true if passed date is before current one.
// Otherwise, and also if dates are equal, returns false.
func (d Date) Before(e Date) bool {
	if d.year < e.year {
		return true
	}
	if d.year > e.year {
		return false
	}
	if d.month < e.month {
		return true
	}
	if d.month > e.month {
		return false
	}
	if d.day < e.day {
		return true
	}
	return false
}

// Sub subtracts passed date and returns duration between them.
// Returned duration is value between their midnights.
func (d Date) Sub(e Date) time.Duration {
	return d.Time().Sub(e.Time())
}

// DaysBetween returns count of days between passed date and current one.
// Result is based on value between their midnights and can be affected by maximum under-laying data type values.
func (d Date) DaysBetween(e Date) int {
	return int(d.Time().Sub(e.Time()).Hours() / 24)
}

// MarshalBinary converts date to binary representation.
// It never returns error.
//
// Byte positions:
//   0       1-4  5     6
//   version year month day
func (d Date) MarshalBinary() ([]byte, error) {
	y := d.year + 1
	return []byte{
		version,
		byte(y >> 24),
		byte(y >> 16),
		byte(y >> 8),
		byte(y),
		d.month + 1,
		d.day + 1,
	}, nil
}

// UnmarshalBinary sets date from passed data.
// It can return wrapped ErrUnsupportedVersion or ErrInvalidLength.
func (d *Date) UnmarshalBinary(data []byte) error {
	l := len(data)
	if l == 0 {
		return fmt.Errorf("Date.UnmarshalBinary: %w: empty data", ErrInvalidLength)
	}
	if data[0] != version {
		return fmt.Errorf("Date.UnmarshalBinary: %w: expected %d instead of %d)", ErrUnsupportedVersion, version, data[0])
	}
	if l != 7 { // version(1)+year(4)+month(1)+day(1)
		return fmt.Errorf("Date.UnmarshalBinary: %w: expected 7 instead of %d", ErrInvalidLength, l)
	}
	d.year = (int32(data[1])<<24 | int32(data[2])<<16 | int32(data[3])<<8 | int32(data[4])) - 1
	d.month = data[5] - 1
	d.day = data[6] - 1
	return nil
}

// MarshalText converts date to text with Formatter.
func (d Date) MarshalText() ([]byte, error) {
	return Formatter(nil, d, 0)
}

// UnmarshalText using global Parser function.
func (d *Date) UnmarshalText(data []byte) error {
	date, err := Parser(data, 0)
	if err != nil {
		return err
	}
	*d = date
	return nil
}

// Scan is support for database/sql package.
// It can return wrapped ErrInvalidType.
func (d *Date) Scan(src any) error {
	if t, ok := src.(time.Time); ok {
		d.FromTime(t)
		return nil
	}
	return fmt.Errorf("%w: expected time.Time instead of %T", ErrInvalidType, src)
}

// Value is support for database/sql package.
func (d Date) Value() (driver.Value, error) {
	return d.Time(), nil
}

// String formats date for string output.
// If Formatter returns error, String returns same value as DefaultFormatter.
func (d Date) String() string {
	b, err := Formatter(nil, d, 0)
	if err != nil {
		b, _ = DefaultFormatter(nil, d, 0)
	}
	return string(b)
}
