// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FilterFromTo_invalid(t *testing.T) {
	date1 := New(2020, August, 7)
	date2 := New(2020, August, 5)
	dateFilter, err := FilterFromTo(&date1, &date2)
	assert.Nil(t, dateFilter)
	assert.EqualError(t, err, "date.FilterFromTo: invalid from or to: 2020-08-07 > 2020-08-05")
}

func Test_FilterFromTo_nil_nil(t *testing.T) {
	dateFilter, err := FilterFromTo(nil, nil)
	require.NoError(t, err)
	assert.True(t, dateFilter.Contains(Date{}))
	assert.True(t, dateFilter.Contains(New(2020, August, 7)))
}

func Test_FilterFromTo_same_date(t *testing.T) {
	date := New(2020, August, 7)
	dateFilter, err := FilterFromTo(&date, &date)
	require.NoError(t, err)
	assert.False(t, dateFilter.Contains(Date{}))
	assert.False(t, dateFilter.Contains(New(2020, August, 6)))
	assert.True(t, dateFilter.Contains(New(2020, August, 7)))
	assert.False(t, dateFilter.Contains(New(2020, August, 8)))
}

func Test_FilterFromTo_date_nil(t *testing.T) {
	date := New(2020, August, 7)
	dateFilter, err := FilterFromTo(&date, nil)
	require.NoError(t, err)
	assert.False(t, dateFilter.Contains(Date{}))
	assert.False(t, dateFilter.Contains(New(2020, August, 6)))
	assert.True(t, dateFilter.Contains(New(2020, August, 7)))
	assert.True(t, dateFilter.Contains(New(2020, August, 8)))
}

func Test_FilterFromTo_nil_date(t *testing.T) {
	date := New(2020, August, 7)
	dateFilter, err := FilterFromTo(nil, &date)
	require.NoError(t, err)
	assert.True(t, dateFilter.Contains(Date{}))
	assert.True(t, dateFilter.Contains(New(2020, August, 6)))
	assert.True(t, dateFilter.Contains(New(2020, August, 7)))
	assert.False(t, dateFilter.Contains(New(2020, August, 8)))
}

func Test_FilterFromTo_date_date(t *testing.T) {
	date1 := New(2020, August, 5)
	date2 := New(2020, August, 7)
	dateFilter, err := FilterFromTo(&date1, &date2)
	require.NoError(t, err)
	assert.False(t, dateFilter.Contains(Date{}))
	assert.False(t, dateFilter.Contains(New(2020, August, 4)))
	assert.True(t, dateFilter.Contains(New(2020, August, 5)))
	assert.True(t, dateFilter.Contains(New(2020, August, 6)))
	assert.True(t, dateFilter.Contains(New(2020, August, 7)))
	assert.False(t, dateFilter.Contains(New(2020, August, 8)))
}

func ExampleFilterFromTo() {
	from := New(2020, August, 5)
	to := New(2020, August, 7)
	dateFilter, _ := FilterFromTo(&from, &to)
	fmt.Println("2020-08-04", dateFilter.Contains(New(2020, August, 4)))
	fmt.Println("2020-08-05", dateFilter.Contains(New(2020, August, 5)))
	fmt.Println("2020-08-06", dateFilter.Contains(New(2020, August, 6)))
	fmt.Println("2020-08-07", dateFilter.Contains(New(2020, August, 7)))
	fmt.Println("2020-08-08", dateFilter.Contains(New(2020, August, 8)))
	// Output:
	// 2020-08-04 false
	// 2020-08-05 true
	// 2020-08-06 true
	// 2020-08-07 true
	// 2020-08-08 false
}
