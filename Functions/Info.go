package Modify

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRoom(farm *AntFarm, line string, isStart, isEnd bool) error {
	parts := strings.Fields(line) //splits the line by each whitespace
	if len(parts) != 3 {          //if it isnt 3, it means it is not a Room
		return fmt.Errorf("invalid room format: %s", line)
	}
	//otherwise the rooms name it in parts[0]
	name := parts[0]
	//the x coord is the in parts[1]
	x, err := strconv.Atoi(parts[1]) //checks if its a valid number
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %s", parts[1])
	}
	//the y coord is the in parts[2]
	y, err := strconv.Atoi(parts[2]) //checks if its a valid number
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %s", parts[2])
	}

	//if the name is duplicated
	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("duplicate room: %s", name) //an error is returned
	}

	//otherwise all info is added to the farm(*AntFarm) struct
	farm.Rooms[name] = &Room{Name: name, X: x, Y: y}

	//to know where to start
	if isStart {
		farm.Start = name //if it was from ##start
	}
	if isEnd {
		farm.End = name //if it was from ##end
	}
	//and where to end

	return nil //doesnt return anything if no errors were found
}

func ParseLink(farm *AntFarm, line string) error {
	parts := strings.Split(line, "-") //splits by -
	if len(parts) != 2 {              //if the length of the array is not 2. it means it is not a link
		return fmt.Errorf("invalid link format: %s", line) //returns an error
	}

	room1, room2 := parts[0], parts[1] // room1 = parts[0] && room2 = parts[1]

	if _, exists := farm.Rooms[room1]; !exists { //checks if the room already exists in the struct
		//if not then it is an unknown room
		return fmt.Errorf("unknown room in link: %s", room1) //an error is then returned
	}
	//same thing for the other room
	if _, exists := farm.Rooms[room2]; !exists {
		return fmt.Errorf("unknown room in link: %s", room2)
	}

	// if they were both valid then they are added to the links part of the struct
	farm.Links[room1] = append(farm.Links[room1], room2)
	farm.Links[room2] = append(farm.Links[room2], room1)

	return nil //doesnt return anything if no errors were found
}

func FindShortestPath(farm *AntFarm) ([]string, error) {
	queue := []string{farm.Start}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == farm.End {
			return ReconstructPath(parent, farm.Start, farm.End), nil
		}

		for _, neighbor := range farm.Links[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, fmt.Errorf("no path found from start to end")
}

func ReconstructPath(parent map[string]string, start, end string) []string {
	path := []string{end}
	for current := end; current != start; current = parent[current] {
		path = append([]string{parent[current]}, path...)
	}
	return path
}

func MoveAnts(farm *AntFarm, path []string) {
	antPositions := make([]int, farm.AntCount)
	for i := range antPositions {
		antPositions[i] = -1 // -1 means the ant hasn't started moving yet
	}

	for {
		moved := false
		movements := make([]string, 0)

		for ant := 0; ant < farm.AntCount; ant++ {
			if antPositions[ant] < len(path)-1 {
				nextPos := antPositions[ant] + 1
				if nextPos == 0 || path[nextPos] == farm.End || IsRoomEmpty(antPositions, nextPos, path) {
					antPositions[ant] = nextPos
					moved = true
					// Only add movement if the ant is not in the start room
					if nextPos > 0 {
						movements = append(movements, fmt.Sprintf("L%d-%s", ant+1, path[nextPos]))
					}
				}
			}
		}

		if !moved {
			break
		}

		// Only print movements if there are any
		if len(movements) > 0 {
			fmt.Println(strings.Join(movements, " "))
		}
	}
}

func IsRoomEmpty(antPositions []int, position int, path []string) bool {
	for _, pos := range antPositions {
		if pos == position {
			return false
		}
	}
	return true
}
