package Modify

import (
	"regexp"
	"strings"
)

func All(s string) string {
	Full := NumOfAnts(s) + "\n" + SCoord(s) + "\n" + ECoord(s) + "\n" + Links(s)
	Full = strings.TrimSpace(Full)
	return Full
}

func NumOfAnts(s string) string {
	s = strings.TrimSpace(s)
	var str string
	re := regexp.MustCompile(`^(\d+)\s*##start`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		str = "Number of Ants: " + matches[1] // Return only the number of ants
	}
	return str
}

func SCoord(s string) string {
	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`[t]\s((?:(\d|[a-z])\s*\d+\s\d\s)+)`)
	numMatches := num.FindStringSubmatch(s)
	var str string
	if len(numMatches) > 0 {

		str = "\nStart co-ords are:\n" + numMatches[1]

	}
	return str
}

func ECoord(s string) string {
	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`[d]\s((?:(\d|[a-z])\s*\d+\s\d\s)+)`)
	numMatches := num.FindStringSubmatch(s)
	var str string
	if len(numMatches) > 0 {

		str = "\nEnd co-ords are: " + numMatches[1]

	}
	return str
}

func Links(s string) string {
	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`\d-\d`)
	numMatches := num.FindAllString(s, -1)
	var str string
	if len(numMatches) > 0 {
		for i := 0; i < len(numMatches); i++ {
			str += "Links are: " + numMatches[i] + "\n"
		}
	}
	return str
}
