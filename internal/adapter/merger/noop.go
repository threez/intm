package merger

import (
	"github.com/threez/intm/internal/model"
	"github.com/threez/intm/internal/port"
)

var _ port.Merger = (*Noop)(nil)

type Noop struct {
	intervals []*model.Interval
}

func (m *Noop) MergeInterval(i *model.Interval) {
	m.intervals = append(m.intervals, i)
}

func (m *Noop) Result() []*model.Interval {
	return m.intervals
}
