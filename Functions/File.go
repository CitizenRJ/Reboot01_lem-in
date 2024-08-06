package Modify

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename) //opens the file
	if err != nil {
		return nil, err
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms: make(map[string]*Room),
		Links: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file) //scans the file line bt line
	lineNum := 0
	parsingLinks := false

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++ //+1 because it scanned a line

		if lineNum == 1 { //first line has to be a valid ant number
			//checks if the first line for the ant number
			farm.AntCount, err = strconv.Atoi(line)
			//if ant number is less than 0 or there is an invalid character it returns an error
			if err != nil || farm.AntCount <= 0 {
				return nil, fmt.Errorf("invalid number of ants")
			}
			continue //if a valid number was found it continues to the next line
		}

		if strings.HasPrefix(line, "#") { //looks for the character #
			// if it was found
			if line == "##start" { //it checks if it the start
				if !scanner.Scan() { //if none of them was a start room
					return nil, fmt.Errorf("missing start room") //an error is returned
				}
				//if it was found line num is then increased
				lineNum++
				// it checks if the lines after it is a Room
				if err := ParseRoom(farm, scanner.Text(), true, false); err != nil {
					return nil, err
				}
			} else if line == "##end" { //it checks if it the end
				if !scanner.Scan() { //if none of them was an end room
					return nil, fmt.Errorf("missing end room") //an error is returned
				}
				//if it was found line num is then increased
				lineNum++
				// it checks if the lines after it is a Room
				if err := ParseRoom(farm, scanner.Text(), false, true); err != nil {
					return nil, err
				}
			}
			continue
		}

		//if parsingLinks is false
		if !parsingLinks {
			//it checks if the line contains the character - (this is for the links)
			if strings.Contains(line, "-") {
				parsingLinks = true //it then sets parsingLinks to true
			} else { //otherswise it checks if it is a Room
				if err := ParseRoom(farm, line, false, false); err != nil {
					return nil, err
				}
				continue
			}
		}

		//if parsingLinks is true
		if parsingLinks {
			//checks if it is a valid link
			if err := ParseLink(farm, line); err != nil {
				return nil, err
			}
		}
	}

	//if no ##start or ##end were found it would return an error
	if farm.Start == "" || farm.End == "" {
		return nil, fmt.Errorf("missing start or end room")
	}

	//otherswise it would send farm with its new info in it
	return farm, nil
}
