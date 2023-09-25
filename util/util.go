package util

import "strings"

func AsLines(s string) []string {
	lines := strings.Split(s, "\n")

	// trim trailing blank line (expected)
	if lines[len(lines) - 1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

var Empty struct{}
