package abstractions

type Office struct {
	List      *List
	OtherList *List
}

func NewOffice(
	list *List,
	otherList *List,
) *Office {
	return &Office{list, otherList}
}
