package abstractions

type Update struct {
	Id UpdateId

	pageNumbers []PageNumber
}

func NewUpdate(
	id UpdateId,
	pageNumbers []PageNumber,
) *Update {
	return &Update{id, pageNumbers}
}
