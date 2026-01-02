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

func (u Update) IsOrderValid(
	protocol *SafetyProtocol,
) bool {

	for pageIndex := len(u.pageNumbers) - 1; pageIndex >= 1; pageIndex-- {
		if !protocol.IsValid(u.pageNumbers[pageIndex], u.pageNumbers[:pageIndex]) {
			return false
		}
	}

	return true
}

func (u Update) GetMiddlePageNumber() PageNumber {
	middlePageIndex := int(float64(len(u.pageNumbers)-1) / float64(2))
	return u.pageNumbers[middlePageIndex]
}
