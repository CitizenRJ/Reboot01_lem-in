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

	txt, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	farm, err := Modify.ParseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Print the input
	fmt.Println(string(txt))
	fmt.Println()

	// Find the shortest path
	path, err := Modify.FindShortestPath(farm)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Move ants through the path
	Modify.MoveAnts(farm, path)
}
