package Modify

type Room struct {
	Name string
	X, Y int
}

type AntFarm struct {
	AntCount int
	Rooms    map[string]*Room
	Links    map[string][]string
	Start    string
	End      string
}
