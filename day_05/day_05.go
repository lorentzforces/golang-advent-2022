package day_05

import (
	"errors"
	"fmt"
	"local-advent-2022/util"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	lines := util.AsLines(input)
	gameLines, moveLines := splitLinesOnBlank(lines)

	gameState := parseGameState(gameLines)
	moves, err := parseMoves(moveLines)
	if err != nil {
		panic(err)
	}

	for _, move := range moves {
		gameState.move(move)
	}

	topLetters := []byte("")
	for _, stack := range gameState.stacks {
		c, success := stack.peek()
		if success {
			topLetters = append(topLetters, c)
		}
	}

	return string(topLetters)
}

func PartTwo(input string) string {
	lines := util.AsLines(input)
	gameLines, moveLines := splitLinesOnBlank(lines)

	gameState := parseGameState(gameLines)
	moves, err := parseMoves(moveLines)
	if err != nil {
		panic(err)
	}

	for _, move := range moves {
		gameState.moveSubStack(move)
	}

	topLetters := []byte("")
	for _, stack := range gameState.stacks {
		c, success := stack.peek()
		if success {
			topLetters = append(topLetters, c)
		}
	}

	return string(topLetters)
}

func parseGameState(lines []string) *craneGame {
	// we're ignoring malformed input for this one, mostly because I don't feel like writing an error- robust parser for this format
	columnLine := lines[len(lines) - 1]

	columns := make([]int, 0)
	for i := 0; i < len(columnLine); i++ {
		if columnLine[i] != ' ' {
			columns = append(columns, i)
		}
	}

	gameState := makeWithColumns(len(columns))
	// skip the column label line and iterate from the bottom-up of the game state diagram
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		// grab values from each column index we detected before
		for j, col := range columns {
			if col >= len(line) {
				continue
			}

			char := line[col]
			if char != ' ' {
				gameState.stacks[j].push(char)
			}
		}
	}
	return gameState
}

func parseMoves(lines []string) ([]craneMove, error) {
	moves := make([]craneMove, 0, len(lines))
	for _, line := range lines {
		move, err := parseMove(line)
		if err != nil {
			return nil, err
		}

		moves = append(moves, move)
	}
	return moves, nil
}

func parseMove(line string) (craneMove, error) {
	words := strings.Fields(line)
	if len(words) != 6 {
		err := errors.New(fmt.Sprintf(
			"Expected 6 terms in move, but found %d in line: \"%s\"",
			len(words),
			line,
		))
		return craneMove{}, err
	}
	hasWrongWords := words[0] != "move" || words[2] != "from" || words[4] != "to"
	if hasWrongWords {
		err := errors.New(fmt.Sprintf(
			"Expected a string of form \"move X from Y to Z\", but was given line: \"%s\"",
			line,
		))
		return craneMove{}, err
	}

	count, err := strconv.ParseUint(words[1], 10, 0)
	if err != nil {
		err := errors.New(fmt.Sprintf(
			"Error parsing count numeric value \"%s\" in move string: \"%s\"",
			words[1],
			line,
		))
		return craneMove{}, err
	}
	src, err := strconv.ParseUint(words[3], 10, 0)
	if err != nil {
		err := errors.New(fmt.Sprintf(
			"Error parsing source numeric value \"%s\" in move string: \"%s\"",
			words[3],
			line,
		))
		return craneMove{}, err
	}
	dest, err := strconv.ParseUint(words[5], 10, 0)
	if err != nil {
		err := errors.New(fmt.Sprintf(
			"Error parsing destination numeric value \"%s\" in move string: \"%s\"",
			words[5],
			line,
		))
		return craneMove{}, err
	}

	return craneMove{
		count: uint(count),
		src: uint(src),
		dest: uint(dest),
	},
	nil
}

func splitLinesOnBlank(lines []string) (gameLines []string, moveLines []string) {
	splitLine := 0
	for i, line := range lines {
		if line == "" {
			splitLine = i
			break
		}
	}

	return lines[:splitLine], lines[splitLine + 1:]
}
