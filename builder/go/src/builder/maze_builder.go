package builder

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(int)
	BuildDoor(int, int)

	GetMaze() *Maze
}

type BaseMazeBuilder struct{}

func (bmb *BaseMazeBuilder) BuildMaze()                     {}
func (bmb *BaseMazeBuilder) BuildRoom(id int)               {}
func (bmb *BaseMazeBuilder) BuildDoor(roomFrom, roomTo int) {}
func (bmb *BaseMazeBuilder) GetMaze() *Maze                 { return nil }

type StandardMazeBuilder struct {
	*BaseMazeBuilder
	_currentMaze *Maze
}

func NewStandardMazeBuilder() *StandardMazeBuilder {
	return &StandardMazeBuilder{
		new(BaseMazeBuilder),
		_currentMaze: nil,
	}
}

func (smb *StandardMazeBuilder) BuildMaze() {
	smb._currentMaze = new(Maze)
}

func (smb *StandardMazeBuilder) BuildRoom(id int) {
	if !smb._currentMaze.RoomNo(id) {
		room := NewRoom(n)
		_currentMaze.AddRoom(room)

		room.SetSide(North, new(Wall))
		room.SetSide(South, new(Wall))
		room.SetSide(East, new(Wall))
		room.SetSide(West, new(Wall))
	}
}

func (smb *StandardMazeBuilder) BuildDoor(idFrom, idTo int) {
	roomFrom := smb._currentMaze.GetRoom(idFrom)
	roomTo := smb._currentMaze.GetRoom(idTo)
	door := NewDoor(roomFrom, roomTo)

	roomFrom.SetSide(CommonWall(roomFrom, roomTo), door)
	roomTo.SetSide(CommonWall(roomTo, roomFrom), door)
}

func (smb *StandardMazeBuilder) GetMaze() *Maze {
	return smb._currentMaze
}
