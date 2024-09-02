package Modify

import (
	"fmt"
)

func DistributeAnts(sortedPaths []*Path, antFarm *AntFarm) {
	//will store the path index for each ant.
	antPaths := make(map[int]int)
	//will store the current position of each ant on its assigned path.
	antPositions := make(map[int]int)
	//will keep track of the number of ants assigned to each path.
	antsInPath := make([]int, len(sortedPaths))

	//this for loop is going throght all the ants
	for antID := 1; antID <= antFarm.numberOfAnts; antID++ {
		//we set the pathIndex to 0
		pathIndex := 0
		//the min cost is the the number of ants in the first path + number of rooms in path
		minCost := antsInPath[0] + sortedPaths[0].numberOfRooms
		for i := 1; i < len(sortedPaths); i++ {
			//the cost is the number of rooms in a path + number of ants in the same path
			cost := sortedPaths[i].numberOfRooms + antsInPath[i]
			if minCost > cost { // if the min is greater than the cost
				minCost = cost //the minCost will take the value of the cost, to find the path with the minimum overall cost for each ant.
				pathIndex = i  // to send the ant to this path.
			}
		}
		//will store the antID to the path index which is the shortest path
		antPaths[antID] = pathIndex
		//will increment the path queue to know how many ants in that path
		antsInPath[pathIndex]++
	}

	//This map will keep track of the ants that are currently outside their assigned paths
	antsOutside := make(map[int][]int)

	//This ensures that each path index has an associated empty slice to store the ant IDs waiting to enter that path.
	for i := 0; i < len(sortedPaths); i++ {
		antsOutside[i] = make([]int, 0)
	}

	//For each ant ID i, it appends the ant ID to the slice associated with the corresponding path index in the antsOutside map.
	for i := 1; i <= len(antPaths); i++ {
		antsOutside[antPaths[i]] = append(antsOutside[antPaths[i]], i)
	}

	//This map will keep track of the ants that are currently inside their assigned paths
	antsInside := make(map[int][]int)
	var antMoving bool
	var output string

	//This is the main loop of the simulation, which continues until no more ants are moving.
	for step := 1; ; step++ {
		antMoving = false
		//This loop iterates through each path index.
		for pathIndex := 0; pathIndex < len(sortedPaths); pathIndex++ {
			//checks the ants that are already inside the path
			for j := 0; j < len(antsInside[pathIndex]); j++ {
				//if the ant has not reached the last room, the ant is moved forward one room
				if antPositions[antsInside[pathIndex][j]] < sortedPaths[pathIndex].numberOfRooms-1 {
					//The antMoving flag is set to true if any ant moves during this step.
					antMoving = true
					antPositions[antsInside[pathIndex][j]]++
					output += "L"
					output += fmt.Sprint(antsInside[pathIndex][j])
					output += "-"
					output += sortedPaths[pathIndex].rooms[antPositions[antsInside[pathIndex][j]]].name
					output += " "

				}
			}
		}
		//This loop iterates through each path index.
		for pathIndex := 0; pathIndex < len(sortedPaths); pathIndex++ {
			//For each path index, it checks the ants that are waiting outside the path
			for len(antsOutside[pathIndex]) != 0 {
				//If the first ant in the antsOutside[pathIndex] slice can move forward one room, the ant is moved
				if antPositions[antsOutside[pathIndex][0]] < sortedPaths[pathIndex].numberOfRooms-1 {
					//The antMoving flag is set to true if any ant moves during this step.
					antMoving = true
					antPositions[antsOutside[pathIndex][0]]++
					output += "L"
					output += fmt.Sprint(antsOutside[pathIndex][0])
					output += "-"
					output += sortedPaths[pathIndex].rooms[antPositions[antsOutside[pathIndex][0]]].name
					output += " "
					antID := antsOutside[pathIndex][0]
					//the ant is removed from the antsOutside slice and added to the antsInside slice
					antsOutside[pathIndex] = antsOutside[pathIndex][1:]
					antsInside[pathIndex] = append(antsInside[pathIndex], antID)
					break // we just want to pass one ant at a time
				}
			}
		}
		//no ants moved during the current step, meaning there are no more ants to be moved
		if !antMoving {
			break
		}
		//The current step number and the accumulated output for that step are printed
		fmt.Println(step, output)
		//empty, to store the next step and output.
		output = ""
	}
}
