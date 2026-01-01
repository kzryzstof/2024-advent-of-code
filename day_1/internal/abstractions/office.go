package abstractions

type Office struct {
	List      *SortedLocationIds
	OtherList *SortedLocationIds
}

func NewOffice(
	list *SortedLocationIds,
	otherList *SortedLocationIds,
) *Office {
	return &Office{list, otherList}
}
