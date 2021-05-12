package merger

import (
	"github.com/google/btree"
	"github.com/threez/intm/internal/model"
)

// Btree attempts so sort and extend to reduce the memory
// consumption without requiring the full list to be in
// memory while processing. Processing is O(n*log(n)),
// Memory is O(n)
type Btree struct {
	tree *btree.BTree
}

func NewBtree() *Btree {
	return &Btree{
		tree: btree.New(2),
	}
}

func (m *Btree) MergeInterval(i *model.Interval) {
	item := &Item{i}

	removed := m.tree.ReplaceOrInsert(item)
	if removed != nil {
		// the interval is overlapping, extend  with the
		// replaced value
		item.Interval.Extend(removed.(*Item).Interval)
	}

	var markedItems []btree.Item
	// function will check if there is an overlapping
	// item and then extend itself, marking all extended
	// items for deletion
	fn := func(i btree.Item) bool {
		// skip self
		if item == i {
			return true // continue search
		}

		ival := i.(*Item).Interval
		if ival.HasOverlap(item.Interval) {
			ival.Extend(item.Interval)
			// mark extend items for later deletion
			markedItems = append(markedItems, item)
			return true // continue search
		}
		return false // stop search
	}

	// search for items that need to be deleted
	// left and right from the current item
	m.tree.DescendLessOrEqual(item, fn)
	m.tree.AscendGreaterOrEqual(item, fn)

	for _, mi := range markedItems {
		m.tree.Delete(mi)
	}
}

func (m *Btree) Result() []*model.Interval {
	var list []*model.Interval
	m.tree.Ascend(func(i btree.Item) bool {
		it := i.(*Item)
		list = append(list, it.Interval)
		return true
	})
	m.tree.Clear(false)
	return list
}

type Item struct {
	*model.Interval
}

func (i Item) Less(bi btree.Item) bool {
	return i.Interval.Before(bi.(*Item).Interval)
}
