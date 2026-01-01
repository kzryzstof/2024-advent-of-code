package abstractions

type Grid struct {
	letters [][]int32
}

func NewGrid(
	letters [][]int32,
) *Grid {
	return &Grid{letters}
}
