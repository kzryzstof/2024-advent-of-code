package abstractions

type Updates struct {
	updates []*Update
}

func NewUpdates(
	updates []*Update,
) *Updates {
	return &Updates{updates}
}

func (u *Updates) CheckUpdates(
	protocol *SafetyProtocol,
) uint64 {

	middlePagesSum := uint64(0)

	for _, update := range u.updates {

		if !update.IsOrderValid(protocol) {
			continue
		}

		middlePagesSum += uint64(update.GetMiddlePageNumber())
	}

	return middlePagesSum
}
