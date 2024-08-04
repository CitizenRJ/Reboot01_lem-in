package main

import (
	Modify "Modify/Functions" // Import the Modify package
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Process the file contents
	content := string(b)

	send := Modify.All(content)
	fmt.Println(send)
}
