package Modify

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRoom(farm *AntFarm, line string, isStart, isEnd bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid room format: %s", line)
	}

	name := parts[0]
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %s", parts[1])
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %s", parts[2])
	}

	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("duplicate room: %s", name)
	}

	farm.Rooms[name] = &Room{Name: name, X: x, Y: y}

	if isStart {
		farm.Start = name
	}
	if isEnd {
		farm.End = name
	}

	return nil
}

func ParseLink(farm *AntFarm, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid link format: %s", line)
	}

	room1, room2 := parts[0], parts[1]

	if _, exists := farm.Rooms[room1]; !exists {
		return fmt.Errorf("unknown room in link: %s", room1)
	}
	if _, exists := farm.Rooms[room2]; !exists {
		return fmt.Errorf("unknown room in link: %s", room2)
	}

	farm.Links[room1] = append(farm.Links[room1], room2)
	farm.Links[room2] = append(farm.Links[room2], room1)

	return nil
}

func PrintInput(farm *AntFarm) {
	fmt.Println(farm.AntCount)
	for _, room := range farm.Rooms {
		prefix := ""
		if room.Name == farm.Start {
			prefix = "##start\n"
		} else if room.Name == farm.End {
			prefix = "##end\n"
		}
		fmt.Printf("%s%s %d %d\n", prefix, room.Name, room.X, room.Y)
	}
	for room, links := range farm.Links {
		for _, link := range links {
			if room < link {
				fmt.Printf("%s-%s\n", room, link)
			}
		}
	}
	fmt.Println()
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
