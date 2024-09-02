package Modify

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filepath string) (*AntFarm, error) {
	// Opening File
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err // Return error if file opening doesnt work
	}
	defer file.Close() // Close file when function is closed

	// Initializing Scanner and AntFarm structure
	fileE, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err // Return error if file opening doesnt work
	}
	
	scanner := bufio.NewScanner(file) // Scanner object created to read file line by line
	antFarm := &AntFarm{
		rooms: make(map[string]*Room),
	}

	if Compare(string(fileE)) {
		antFarm.edgeCase = true
	}

	// Storing either start or end room
	var pendingType string
	var startcount, endcount int
	// Reading file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim the line for any trailing whitespace

		if line == "" { // If no contents in the line
			continue
		}

		if antFarm.numberOfAnts == 0 {
			antFarm.numberOfAnts, err = strconv.Atoi(line) //gets the ant number
			if err != nil  || antFarm.numberOfAnts <= 0 || antFarm.numberOfAnts > 1000{
				return nil, fmt.Errorf("ERROR: invalid data format, invalid ant number")
			}
			continue
		}
		
		if strings.HasPrefix(line, "#") { // If the line has a prefix with # we check the Room type either start or end
			if line == "##start" {
				pendingType = "start"
				startcount++
			} else if line == "##end" {
				pendingType = "end"
				endcount++
			}
			
			continue // We skip if we have comments or anything that is not a starting or ending room
		}
		
		// Parsing Rooms
		if strings.Contains(line, " ") {
			parts := strings.Split(line, " ") // []string
			if len(parts) != 3 {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room format")
			}

			// Parse to store coordinates for bonus repository
			x, _ := strconv.Atoi(parts[1]) // Coordinate X
			y, _ := strconv.Atoi(parts[2]) // Coordinate Y

			room := &Room{
				name:  parts[0],
				x:     x,
				y:     y,
				links: []*Room{},
			}

			antFarm.rooms[room.name] = room // The room name identifies the room struct for O(1) lookup

			if pendingType == "start" {
				antFarm.startRoom = room // If the pendingType is populated with "start", then the room after that line directly is the start room
				pendingType = ""
			} else if pendingType == "end" { // If the pendingType is populated with "end", then the room after that line directly is the end room
				antFarm.endRoom = room
				pendingType = ""
			}
		}

		// Parsing Links
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid link format") // ERROR: invalid data format, Invalid link format
			}
			if _, exists := antFarm.rooms[parts[0]]; !exists {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room name")
			}
			if _, exists := antFarm.rooms[parts[1]]; !exists {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room name")
			}
			room1 := antFarm.rooms[parts[0]]
			room2 := antFarm.rooms[parts[1]]
			room1.links = append(room1.links, room2)
			room2.links = append(room2.links, room1)
		}
	}

	// Error: Hits an end of file condition without reading any data
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("scanner Error")
	}

	// Error: Ant Farm does not have a start or end room

	if antFarm.startRoom == nil || antFarm.endRoom == nil {
		return nil, fmt.Errorf("ERROR: invalid data format, missing Start Room or End Room")
	}
	if startcount > 1 || endcount>1{
		return nil, fmt.Errorf("ERROR: invalid data format, invalid number of start/end")
	}

	return antFarm, nil
}
