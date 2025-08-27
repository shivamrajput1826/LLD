package game

type Position struct {
	FinalPos int
}

type Grid struct {
	SnakePos   *Position
	LadderPos  *Position
	CurrentPos int
}

type Board struct {
	Grid [100]*Grid
}

func NewBoard() *Board {
	board := Board{}
	board.Initialize()
	board.InitializeSnakes()
	board.InitializeLadders()
	return &board
}

func (b *Board) Initialize() {
	for i := 0; i < 100; i++ {
		b.Grid[i] = &Grid{
			CurrentPos: i,
		}
	}
}

func (b *Board) InitializeSnakes() {
	b.Grid[55].SnakePos = &Position{FinalPos: 10}

	b.Grid[98].SnakePos = &Position{FinalPos: 7}

	b.Grid[70].SnakePos = &Position{FinalPos: 33}

	b.Grid[47].SnakePos = &Position{FinalPos: 15}

	b.Grid[62].SnakePos = &Position{FinalPos: 19}
}

func (b *Board) InitializeLadders() {
	b.Grid[3].LadderPos = &Position{FinalPos: 22}

	b.Grid[8].LadderPos = &Position{FinalPos: 30}

	b.Grid[20].LadderPos = &Position{FinalPos: 41}

	b.Grid[28].LadderPos = &Position{FinalPos: 84}

	b.Grid[58].LadderPos = &Position{FinalPos: 77}
}

func (b *Board) GetNewPosition(pos, value int) int {
	newPos := pos + value
	if newPos > 99 {
		return pos
	}
	cell := b.Grid[newPos]
	if cell.SnakePos != nil {
		return cell.SnakePos.FinalPos
	} else if cell.LadderPos != nil {
		return cell.LadderPos.FinalPos
	}
	return newPos
}
