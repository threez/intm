package reader

import (
	"bufio"
	"fmt"
	"io"

	"github.com/threez/intm/internal/model"
	"github.com/threez/intm/internal/port"
)

var _ port.IntervalReader = (*Reader)(nil)

type Reader struct {
	r    *bufio.Reader
	line int
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

func (r *Reader) Read() (*model.Interval, error) {
	line, _, err := r.r.ReadLine()
	r.line++
	if err != nil {
		return nil, err
	}

	var i model.Interval
	n, err := fmt.Sscanf(string(line), "%d %d", &i.Start, &i.End)
	// if n != 2 than scanning failed,
	// provide to user info about failed input
	if err != nil {
		return nil, fmt.Errorf("%w in line %d: %v", port.ErrInvalidContent, r.line, err)
	}
	if n != 2 {
		return nil, fmt.Errorf("%w in line %d: invalid format, expected two numbers", port.ErrInvalidContent, r.line)
	}

	return &i, nil
}
