package port

import "github.com/threez/intm/internal/model"

// Merger represents a merge interval implementation
type Merger interface {
	// MergeInterval merges the passed interval
	MergeInterval(*model.Interval)
	// Result returns the list of all merged intervals
	// including all overlapping intervals of the
	// previous calls to MergeInterval.
	// The function will clear the data. Calling the function
	// twice will result in no data the second time.
	Result() []*model.Interval
}
