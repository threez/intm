package port

import (
	"errors"

	"github.com/threez/intm/internal/model"
)

var ErrInvalidContent = errors.New("invalid interval")

type IntervalReader interface {
	// Read reads an interval at a time and returns
	// nil, io.EOF when there are no more intervals
	Read() (*model.Interval, error)
}
