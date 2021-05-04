package model

type List []*Interval

func (list List) Len() int {
	return len(list)
}

func (list List) Less(i, j int) bool {
	return list[i].Before(list[j])
}

func (list List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
