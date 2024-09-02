package Modify

type Room struct {
	name  string  // Room Identifier (Could be digit/string/whatever)
	x, y  int     // Coordinates for visualization
	links []*Room // Rooms linking to this room struct
}

type Path struct {
	rooms         []*Room // List of rooms in this path
	numberOfRooms int     // Number of rooms
}

type AntFarm struct {
	rooms              map[string]*Room // If visited already, then ignore on second pass of bfs exploration
	numberOfAnts       int              // Number of ants in the ant farm
	startRoom, endRoom *Room            //Stores the start and end room
	edgeCase           bool
}
