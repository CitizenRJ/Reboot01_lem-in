package Modify

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms: make(map[string]*Room),
		Links: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	parsingLinks := false

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if lineNum == 1 {
			farm.AntCount, err = strconv.Atoi(line)
			if err != nil || farm.AntCount <= 0 {
				return nil, fmt.Errorf("invalid number of ants")
			}
			continue
		}

		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				if !scanner.Scan() {
					return nil, fmt.Errorf("missing start room")
				}
				lineNum++
				if err := ParseRoom(farm, scanner.Text(), true, false); err != nil {
					return nil, err
				}
			} else if line == "##end" {
				if !scanner.Scan() {
					return nil, fmt.Errorf("missing end room")
				}
				lineNum++
				if err := ParseRoom(farm, scanner.Text(), false, true); err != nil {
					return nil, err
				}
			}
			continue
		}

		if !parsingLinks {
			if strings.Contains(line, "-") {
				parsingLinks = true
			} else {
				if err := ParseRoom(farm, line, false, false); err != nil {
					return nil, err
				}
				continue
			}
		}

		if parsingLinks {
			if err := ParseLink(farm, line); err != nil {
				return nil, err
			}
		}
	}

	if farm.Start == "" || farm.End == "" {
		return nil, fmt.Errorf("missing start or end room")
	}

	return farm, nil
}
