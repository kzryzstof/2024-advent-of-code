package abstractions

type Grid struct {
	letters [][]int32
}

var (
	Right      = Vector{0, 1}
	UpRight    = Vector{-1, 1}
	Up         = Vector{-1, 0}
	UpLeft     = Vector{-1, -1}
	Left       = Vector{0, -1}
	DownLeft   = Vector{1, -1}
	Down       = Vector{1, 0}
	DownRight  = Vector{1, 1}
	Directions = []Vector{UpRight, Up, UpLeft, Left, DownLeft, Down, DownRight, Right}
)

func NewGrid(
	letters [][]int32,
) *Grid {
	return &Grid{letters}
}

func (g *Grid) CountWord(
	word string,
) int64 {

	count := int64(0)

	for rowIndex, row := range g.letters {
		for columnIndex, letter := range row {
			if letter == 'X' {
				for _, direction := range Directions {
					if g.IsWordPresent(word, rowIndex, columnIndex, direction) {
						count++
					}
				}
			}
		}
	}

	return count
}

func (g *Grid) IsWordPresent(
	word string,
	fromRow int,
	fromCol int,
	direction Vector,
) bool {
	for letterIndex := 1; letterIndex < len(word); letterIndex++ {
		row := fromRow + letterIndex*int(direction.X)
		col := fromCol + letterIndex*int(direction.Y)

		if row < 0 || col < 0 || row >= len(g.letters) || col >= len(g.letters[row]) {
			return false
		}

		if g.letters[row][col] != int32(word[letterIndex]) {
			return false
		}
	}

	return true
}
