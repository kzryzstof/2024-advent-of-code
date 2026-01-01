package abstractions

type Grid struct {
	letters [][]int32
}

var (
	UpRight    = Vector{-1, 1}
	UpLeft     = Vector{-1, -1}
	DownLeft   = Vector{1, -1}
	DownRight  = Vector{1, 1}
	Directions = []Vector{
		UpLeft,
		UpRight,
		DownRight,
		DownLeft,
	}

	Patterns = []string{
		"MSSM",
		"MMSS",
		"SMMS",
		"SSMM",
	}
)

func NewGrid(
	letters [][]int32,
) *Grid {
	return &Grid{letters}
}

func (g *Grid) CountXmasPatterns() int64 {

	patternsCount := int64(0)

	for rowIndex, row := range g.letters {
		for columnIndex, letter := range row {
			if letter == 'A' {
				for _, pattern := range Patterns {
					if g.isPattern(pattern, rowIndex, columnIndex) {
						patternsCount++
						break
					}
				}
			}
		}
	}

	return patternsCount
}

func (g *Grid) isPattern(
	pattern string,
	fromRow int,
	fromCol int,
) bool {
	for index, letter := range pattern {
		if !g.isLetterPresent(letter, fromRow, fromCol, Directions[index]) {
			return false
		}
	}

	return true
}

func (g *Grid) isLetterPresent(
	letter int32,
	fromRow int,
	fromCol int,
	direction Vector,
) bool {
	row := fromRow + int(direction.X)
	col := fromCol + int(direction.Y)

	if row < 0 || col < 0 || row >= len(g.letters) || col >= len(g.letters[row]) {
		return false
	}

	if g.letters[row][col] != letter {
		return false
	}

	return true
}
