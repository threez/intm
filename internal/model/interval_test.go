package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interval_Validate(t *testing.T) {
	i := Interval{}
	assert.Equal(t, ErrEndIsEqualToStart, i.Validate())

	i.Start = 10
	i.End = 5
	assert.Equal(t, ErrEndBeforeStart, i.Validate())

	i.End = 25
	assert.Nil(t, i.Validate())
}

func Test_Interval_IsOverLapping(t *testing.T) {
	i1 := Interval{Start: 5, End: 10}

	i2 := Interval{Start: 7, End: 8}
	i3 := Interval{Start: 7, End: 15}
	i4 := Interval{Start: 3, End: 7}

	assert.True(t, i1.HasOverlap(&i2))
	assert.True(t, i1.HasOverlap(&i3))
	assert.True(t, i1.HasOverlap(&i4))

	i5 := Interval{Start: 1, End: 5}
	i6 := Interval{Start: 15, End: 55}

	assert.False(t, i1.HasOverlap(&i5))
	assert.False(t, i1.HasOverlap(&i6))
}
