package builder

var (
	North = 1
	South = 2
	East  = 3
	West  = 4
)

type MazeComponent interface{}

type Wall struct{}

type Room struct{}

func NewRoom(id int) *Room {
	return new(Room)
}

func (r *Room) CommonWall(room *Room) int {
	return North
}

func (r *Room) SetSide(direction int, wall MazeComponent) {}

type Maze struct {
	rooms map[int]*Room
}

func (m *Maze) AddRoom(room *Room) {

}

func (m *Maze) GetRoom(id int) *Room {
	return nil
}

func (m *Maze) RoomNo(id int) *Room {
	return m.rooms[id]
}

type Door struct{}

func NewDoor(from, to *Room) *Door {
	return new(Door)
}
