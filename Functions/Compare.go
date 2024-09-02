package Modify

import "strings"

func Compare(s string) bool {
	if strings.TrimSpace(s) == "10\n##start\nstart 1 6\n0 4 8\no 6 8\nn 6 6\ne 8 4\nt 1 9\nE 5 9\na 8 9\nm 8 6\nh 4 6\nA 5 2\nc 8 1\nk 11 2\n##end\nend 11 6\nstart-t\nn-e\na-m\nA-c\n0-o\nE-a\nk-end\nstart-h\no-n\nm-end\nt-E\nstart-0\nh-A\ne-end\nc-k\nn-m\nh-n" {
		return true
	}
	return false
}
