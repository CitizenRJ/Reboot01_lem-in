package main

import (
	Modify "Modify/Functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	farm, err := Modify.ParseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Print the input
	Modify.PrintInput(farm)

	// Find the shortest path
	path, err := Modify.FindShortestPath(farm)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Move ants through the path
	Modify.MoveAnts(farm, path)
}
