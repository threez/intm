package merger

import (
	"sort"

	"github.com/threez/intm/internal/model"
)

// Simple ignores memory and cpu constraints.
// Alg should be O(n*log(n)) due the use of search
// search. Memory can be max O(2n) since the new list is
// created next to the existing list.
type Simple struct {
	intervals []*model.Interval
}

func NewSimple() *Simple {
	return &Simple{}
}

func (m *Simple) MergeInterval(i *model.Interval) {
	m.intervals = append(m.intervals, i)
}

func (m *Simple) Result() []*model.Interval {
	// Step 1. sort the input
	sort.Sort(model.List(m.intervals))

	// Step 2. extend ranges until not longer possible
	var mergedList []*model.Interval
	for i := 0; i < len(m.intervals); i++ {
		mergedList = append(mergedList, m.intervals[i])
		extends := 0
		for j := i + 1; j < len(m.intervals); j++ {
			if m.intervals[i].HasOverlap(m.intervals[j]) {
				m.intervals[i].Extend(m.intervals[j])
				extends++
			} else {
				// max extend reached, give up search (inner loop)
				break
			}
		}
		i += extends // move i to skip over already extended sections
	}

	m.intervals = nil
	return mergedList
}
