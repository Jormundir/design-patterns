package builder

type MazeGame struct {
	maze *Maze
}

func NewMazeGame() *MazeGame {
	builder := new(StandardMazeBuilder)
	mazeGame := new(MazeGame)

	mazeGame.CreateMaze(builder)
	mazeGame.maze = builder.GetMaze()
	return mazeGame
}

func (mg *MazeGame) CreateMaze(builder MazeBuilder) *Maze {
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

func (mg *MazeGame) CreateComplexMaze(builder MazeBuilder) *Maze {
	builder.BuildRoom(1)
	// ...
	builder.BuildRoom(1000)

	return builder.GetMaze()
}
