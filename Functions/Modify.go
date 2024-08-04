package Modify

import (
	"regexp"
	"strings"
)

func All(s string) string {
	Full := NumOfAnts(s) + "\n"
	strB := Between(s)
	strA := After(s)
	Full += "\nSTART"
	linesB := strings.Split(strings.TrimSpace(strB), "\n")
	for i := 0; i < len(linesB); i++ {
		Full += SCoord(linesB[i])
	}
	Full += "\n\nEND"
	Full += ECoord(strA)
	linesA := strings.Split(strings.TrimSpace(strA), "\n")
	for i := 0; i < len(linesA); i++ {
		Full += Links(linesA[i])
	}
	Full = strings.TrimSpace(Full)
	return Full
}

func NumOfAnts(s string) string {
	s = strings.TrimSpace(s)
	var str string
	re := regexp.MustCompile(`(\d+)\s*##start`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		str = "Number of Ants: " + matches[1] // Return only the number of ants
	}
	return str
}

func Between(s string) string {
	bet := regexp.MustCompile(`(?s)##start\n(.*?)##end`)
	All := bet.FindStringSubmatch(s)[1]
	return All
}

func After(s string) string {
	pattern := `##end`
	regex := regexp.MustCompile(pattern)
	parts := regex.Split(s, -1)
	return parts[1]
}

func SCoord(s string) string {
	//[t]\s((?:(\d|[a-z])\s*\d+\s\d\s)+)
	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`(\d+|[a-z]+)\s(\d+)\s(\d+)`)
	SRoCo := num.FindStringSubmatch(s)
	// fmt.Println(numMatches)
	var str string
	if len(SRoCo) > 0 {
		str = "\nRoom name: " + SRoCo[1] + ", Co-ords are: " + SRoCo[2] + " , " + SRoCo[3]
	}
	return str
}

func ECoord(s string) string {
	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`(\d+|[a-z]+)\s(\d+)\s(\d+)`)
	ERoCo := num.FindStringSubmatch(s)
	var str string
	if len(ERoCo) > 0 {

		str = "\nRoom name: " + ERoCo[1] + " ,Co-ords are: " + ERoCo[2] + " , " + ERoCo[3] + "\n\n"

	}
	return str
}

func Links(s string) string {

	s = strings.TrimSpace(s)
	num := regexp.MustCompile(`(\d+|[a-z]+)-(\d+|[a-z]+)`)
	link := num.FindStringSubmatch(s)
	var str string
	if len(link) > 0 {
		str += "Links are: " + link[1] + "-" + link[2] + "\n"
	}
	return str
}
