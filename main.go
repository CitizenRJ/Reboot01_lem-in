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

	txt, err := os.ReadFile(os.Args[1]) //Reads the txt file
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Print the txt file
	fmt.Println(string(txt))
	fmt.Println()

	/*
		here is checking if there is a ##start
		and a ##end before actually doing anything else
	*/
	farm, err := Modify.ParseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Finds the shortest path
	path, err := Modify.FindShortestPath(farm)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Moves ants through the path
	Modify.MoveAnts(farm, path)
}
