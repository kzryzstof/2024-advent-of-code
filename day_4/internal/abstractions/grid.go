package abstractions

type Grid struct {
	letters [][]int32
}

func NewGrid(
	letters [][]int32,
) *Grid {
	return &Grid{letters}
}

func (g *Grid) CountWord(
	word string,
) int64 {

	for _, row := range g.letters {
		for _, letter := range row {
			if letter == 'X' {
				continue
			}
		}
	}

	return 0
}
