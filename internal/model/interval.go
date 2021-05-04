package model

import (
	"errors"
	"fmt"
)

var (
	ErrEndBeforeStart    = errors.New("end is before start")
	ErrEndIsEqualToStart = errors.New("end is equal to start")
)

type Interval struct {
	Start int
	End   int
}

func NewInterval(start, end int) *Interval {
	return &Interval{
		Start: start,
		End:   end,
	}
}

func (i Interval) String() string {
	return fmt.Sprintf("%d %d", i.Start, i.End)
}

// Validate checks if the Interval is valid
// and returns the first error in case that is not
// the case
func (i Interval) Validate() error {
	if i.Start == i.End {
		return ErrEndIsEqualToStart
	}
	if i.End < i.Start {
		return ErrEndBeforeStart
	}
	return nil
}

// HasOverlap returns true if there is any
// overlap with provided interval
func (i Interval) HasOverlap(o *Interval) bool {
	return (i.Start < o.Start && o.Start <= i.End) ||
		(i.Start < o.End && o.End <= i.End) ||
		(o.Start < i.Start && i.End < o.End)
}

// IsBefore returns true in case the passed interval is after
func (i Interval) Before(o *Interval) bool {
	return i.Start < o.Start
}

// Extend extend the current interval with the provided interval
func (i *Interval) Extend(o *Interval) {
	if o.Start < i.Start {
		i.Start = o.Start
	}
	if o.End > i.End {
		i.End = o.End
	}
}
