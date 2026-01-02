package abstractions

type Updates struct {
	updates []*Update
}

func NewUpdates(
	updates []*Update,
) *Updates {
	return &Updates{updates}
}
