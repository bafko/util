// Copyright 2022 Livesport TV s.r.o. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package date

import (
	"fmt"
)

// Filter represents date filter.
type Filter interface {
	// Contains returns true if passed date is accepted by filter.
	Contains(date Date) bool
}

// FilterFromTo creates new date filter based on from date and to date.
// If from date (or to date) is nil, there isn't from date (or to date) limit.
// From date (or to date) is including to filter.
func FilterFromTo(from, to *Date) (Filter, error) {
	if from == nil {
		if to == nil {
			return filterNo{}, nil
		}
		return &filterTo{
			to: *to,
		}, nil
	}
	if to == nil {
		return &filterFrom{
			from: *from,
		}, nil
	}
	if from.Equal(*to) {
		return &filterDate{
			date: *from,
		}, nil
	}
	if from.After(*to) {
		return nil, fmt.Errorf("%w: %s > %s", ErrInvalidFromOrTo, from, to)
	}
	return &filterFromTo{
		from: *from,
		to:   *to,
	}, nil
}

type filterNo struct {
}

func (filterNo) Contains(_ Date) bool {
	return true
}

type filterDate struct {
	date Date
}

func (d *filterDate) Contains(date Date) bool {
	return d.date.Equal(date)
}

type filterFrom struct {
	from Date
}

func (d *filterFrom) Contains(date Date) bool {
	return d.from.Equal(date) || d.from.Before(date)
}

type filterTo struct {
	to Date
}

func (d *filterTo) Contains(date Date) bool {
	return d.to.Equal(date) || d.to.After(date)
}

type filterFromTo struct {
	from Date
	to   Date
}

func (d *filterFromTo) Contains(date Date) bool {
	return d.from.Equal(date) || d.to.Equal(date) || (d.from.Before(date) && d.to.After(date))
}
