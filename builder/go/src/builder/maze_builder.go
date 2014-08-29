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
		BaseMazeBuilder: new(BaseMazeBuilder),
		_currentMaze:    nil,
	}
}

func (smb *StandardMazeBuilder) BuildDoor(idFrom, idTo int) {
	roomFrom := smb._currentMaze.GetRoom(idFrom)
	roomTo := smb._currentMaze.GetRoom(idTo)
	door := NewDoor(roomFrom, roomTo)

	roomFrom.SetSide(roomFrom.CommonWall(roomTo), door)
	roomTo.SetSide(roomTo.CommonWall(roomFrom), door)
}

func (smb *StandardMazeBuilder) BuildMaze() {
	smb._currentMaze = new(Maze)
}

func (smb *StandardMazeBuilder) BuildRoom(id int) {
	if smb._currentMaze.RoomNo(id) == nil {
		room := NewRoom(id)
		smb._currentMaze.AddRoom(room)

		room.SetSide(North, new(Wall))
		room.SetSide(South, new(Wall))
		room.SetSide(East, new(Wall))
		room.SetSide(West, new(Wall))
	}
}

func (smb *StandardMazeBuilder) GetMaze() *Maze {
	return smb._currentMaze
}
