package merger

import (
	"github.com/threez/intm/internal/model"
)

// Btree attempts so sort and extend to reduce the memory
// consumption without requiring the full list to be in
// memory while processing. Processing is O(n*log(n)),
// Memory is O(n)
type Btree struct {
	root *Node
}

func NewBtree() *Btree {
	return &Btree{}
}

func (m *Btree) MergeInterval(i *model.Interval) {
	// initialize with first element
	if m.root == nil {
		m.root = &Node{Interval: i}
		return
	}

	m.root.MergeInterval(i)
}

func (m *Btree) Result() []*model.Interval {
	if m.root == nil {
		return nil
	}
	var list []*model.Interval
	m.root.Result(&list)
	m.root = nil
	return list
}

type Node struct {
	Before, After *Node
	Interval      *model.Interval
}

func (m *Node) MergeInterval(i *model.Interval) {
	// if the new interval brings an overlap
	// extend the current node
	if m.Interval.HasOverlap(i) {
		m.Interval.Extend(i)

		// check if this node is now overlapping with
		// before or after nodes
		if m.Before != nil && m.Before.Interval.HasOverlap(m.Interval) {
			m.Interval.Extend(m.Before.Interval)
			m.Before = nil
		}

		if m.After != nil && m.After.Interval.HasOverlap(m.Interval) {
			m.Interval.Extend(m.After.Interval)
			m.After = nil
		}

		return
	}

	// else add new node
	if i.Before(m.Interval) {
		if m.Before != nil {
			m.Before.MergeInterval(i)
		} else {
			m.Before = &Node{Interval: i}
		}
	} else {
		if m.After != nil {
			m.After.MergeInterval(i)
		} else {
			m.After = &Node{Interval: i}
		}
	}
}

func (m *Node) Result(list *[]*model.Interval) {
	if m.Before != nil {
		m.Before.Result(list)
	}
	*list = append(*list, m.Interval)
	if m.After != nil {
		m.After.Result(list)
	}
}
