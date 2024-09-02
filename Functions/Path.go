package Modify

import (
	"sort"
)

// The function returns all the shortest paths
func Paths(antFarm *AntFarm) []*Path {

	var paths []*Path
	//contains a single Path with the antFarm.startRoom as the only room in the path, and the number of rooms in the path set to 1.
	queue := []Path{
		{
			rooms:         []*Room{antFarm.startRoom},
			numberOfRooms: 1,
		},
	}
	//We create a map called visited to keep track of the rooms we have already visited. Initially, we mark the antFarm.startRoom as visited.
	visited := map[string]bool{antFarm.startRoom.name: true}

	for len(queue) != 0 {
		//we take the first Path from the queue, remove it from the queue, and get a reference to the last room in the path.
		path := queue[0]
		queue = queue[1:]
		lastRoom := path.rooms[len(path.rooms)-1]
		//If the last room in the current path is the antFarm.endRoom
		if lastRoom == antFarm.endRoom {
			//we reset the visited map to only mark the antFarm.startRoom as visited
			visited = map[string]bool{antFarm.startRoom.name: true}
			//append a pointer to the current path to the paths slice
			paths = append(paths, &path)
			//then update the visited map to mark all the rooms in all the paths found so far as visited
			for _, path := range paths {
				for _, room := range path.rooms {
					visited[room.name] = true
				}
			}

			//we reset the queue to only contain a single path starting from the antFarm.startRoom
			queue = []Path{
				{
					rooms:         []*Room{antFarm.startRoom},
					numberOfRooms: 1,
				},
			}
			//continue to the next loop
			continue
		}
		//If the last room in the current path is not the antFarm.endRoom, we iterate through all the links
		for _, link := range lastRoom.links {
			//For each link, we check if the link has not been visited before,
			//or if the link is the antFarm.endRoom and the last room is not the antFarm.startRoom.
			if !visited[link.name] || (link == antFarm.endRoom && lastRoom != antFarm.startRoom) {
				//we create a new path by copying the rooms from the current path,
				//appending the new link, and adding this new path to the queue.
				newPathRooms := make([]*Room, len(path.rooms))
				copy(newPathRooms, path.rooms)
				newPathRooms = append(newPathRooms, link)
				newPath := Path{
					rooms:         newPathRooms,
					numberOfRooms: len(newPathRooms),
				}
				queue = append(queue, newPath)
				visited[link.name] = true

				if antFarm.edgeCase {
					break
				}
			}
		}
	}

	//we sort the paths slice in ascending order by the numberOfRooms field of each Path struct using an anonymous function,
	sort.SliceStable(paths, func(i, j int) bool {
		return paths[i].numberOfRooms < paths[j].numberOfRooms
	})

	//return the sorted paths slice.
	return paths
}
