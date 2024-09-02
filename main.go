package main

import (
	Modify "Modify/Functions"
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Number of arguments passed in leads to program termination.")
		return
	}
	filepath := os.Args[1]

	// Printing the ant farm line by line
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	antFarm, err := Modify.ParseFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	sortedPaths := Modify.Paths(antFarm)
	if len(sortedPaths) == 0 {
		fmt.Println("ERROR: invalid data format, no link to end")
		return
	}

	for _, line := range fileLines {
		fmt.Println(line)
	}
	fmt.Println()
	
	Modify.DistributeAnts(sortedPaths, antFarm)
}
