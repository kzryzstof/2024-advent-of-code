package abstractions

import "sort"

type SortedLocationIds struct {
	ids []LocationId
}

func NewList(
	ids []LocationId,
) *SortedLocationIds {

	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	return &SortedLocationIds{ids}
}

func (l *SortedLocationIds) Get(
	index uint,
) LocationId {
	return l.ids[index]
}

func (l *SortedLocationIds) Len() uint {
	return uint(len(l.ids))
}
