package main

import (
	"fmt"
)

type MapSite interface {
	Enter()
}
type Wall interface {
	MapSite
}

type Door interface {
	MapSite
	IsOpen() bool
	SetOpen(isOpen bool)
}

type Room interface {
	MapSite
	GetSide(direction string) MapSite
	SetSide(direction string, side MapSite)
	SetRoomId(roomId int)
	GetRoomId() int
}

type MazeFactory interface {
	MakeMaze() *Maze
	MakeWall() Wall
	MakeRoom(roomNo int) Room
	MakeDoor(room1 Room, room2 Room) Door
}

type BombedWall struct {
}

func (w *BombedWall) Enter() {
	//fmt.Println("Can’t enter a wall.")
}

type BombeRoom struct {
	RoomNo int
	Side   map[string]MapSite
}

func (r *BombeRoom) Enter() {
	for _, v := range r.Side {
		v.Enter()
	}
}

func (r *BombeRoom) GetSide(direction string) MapSite {
	return r.Side[direction]
}

func (r *BombeRoom) SetSide(direction string, side MapSite) {
	r.Side[direction] = side
}

func (r *BombeRoom) SetRoomId(roomId int) {
	r.RoomNo = roomId
}

func (r *BombeRoom) GetRoomId() int {
	return r.RoomNo
}

func NewBombedRoom(id int) Room {
	return &BombeRoom{RoomNo: id, Side: make(map[string]MapSite)}
}

type Maze struct {
	rooms map[int]Room
}

type BombedMazeFactory struct{}

func NewMaze() *Maze {
	return &Maze{rooms: make(map[int]Room)}
}

type BombedDoor struct {
	room1  Room
	room2  Room
	isOpen bool
}

func NewBombedDoor(room1 Room, room2 Room) *BombedDoor {
	return &BombedDoor{room1: room1, room2: room2}
}
func (door *BombedDoor) IsOpen() bool {
	return door.isOpen
}
func (door *BombedDoor) SetOpen(isOpen bool) {
	door.isOpen = isOpen
}
func (door *BombedDoor) Enter() {
	if door.isOpen {
		fmt.Println("Enter bombed door")
	} else {
		fmt.Println("Can’t enter bombed door. Closed.")
	}
}
func (door *BombedDoor) String() string {
	return fmt.Sprintf("A bombed door between %v and %v", door.room1.GetRoomId(), door.room2.GetRoomId())
}

func (maze *Maze) AddRoom(room Room) {
	maze.rooms[room.GetRoomId()] = room
}

func (maze *Maze) GetRoom(roomId int) Room {
	for k, v := range maze.rooms {
		if k == roomId {
			return v
		}
	}
	return nil
}

func (factory *BombedMazeFactory) MakeMaze() *Maze {
	return NewMaze()
}

func (factory *BombedMazeFactory) MakeWall() Wall {
	return new(BombedWall)
}

func (factory *BombedMazeFactory) MakeRoom(roomNo int) Room {
	return NewBombedRoom(roomNo)
}

func (factory *BombedMazeFactory) MakeDoor(room1 Room, room2 Room) Door {
	return NewBombedDoor(room1, room2)
}

func CreateMaze(factory MazeFactory) *Maze {
	aMaze := factory.MakeMaze()
	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(room1, room2)

	aDoor.SetOpen(true)

	aMaze.AddRoom(room1)
	aMaze.AddRoom(room2)
	room1.SetSide("North", factory.MakeWall())
	room1.SetSide("East", aDoor)
	room1.SetSide("South", factory.MakeWall())
	room1.SetSide("West", factory.MakeWall())
	room2.SetSide("North", factory.MakeWall())
	room2.SetSide("East", factory.MakeWall())
	room2.SetSide("South", factory.MakeWall())
	room2.SetSide("West", aDoor)
	return aMaze
}

func main() {
	var factory *BombedMazeFactory
	var maze *Maze
	maze = CreateMaze(factory)
	maze.GetRoom(1).Enter() //Prints: Can’t enter bombed door. Closed
}
