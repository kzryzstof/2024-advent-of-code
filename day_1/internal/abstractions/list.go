package abstractions

type List struct {
	ids []LocationId
}

func NewList(
	ids []LocationId,
) *List {
	return &List{ids}
}

func (l *List) Len() uint {
	return uint(len(l.ids))
}
