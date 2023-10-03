package main

import (
	"fmt"
	"local-advent-2022/day_01"
	"local-advent-2022/day_02"
	"local-advent-2022/day_03"
	"local-advent-2022/day_04"
	"local-advent-2022/day_05"
	"os"
	"sort"
)

func main() {
	results := runAll(RUN_DATA)

	sort.Slice(
		results,
		func(i, j int) bool {
			if results[i].day == results[j].day {
				return results[i].part < results[j].part
			}
			return results[i].day < results[j].day
		},
	)

	for _, result := range results {
		fmt.Printf("Day %02d, Part %02d output: %s\n", result.day, result.part, result.output)
	}
}

var RUN_DATA []puzzleData = []puzzleData{
	{
		day: 1,
		part: 1,
		inputFile: "inputs/day_01_input.txt",
		fn: func(input string) any {return day_01.PartOne(input)},
	},
	{
		day: 1,
		part: 2,
		inputFile: "inputs/day_01_input.txt",
		fn: func(input string) any {return day_01.PartTwo(input)},
	},
	{
		day: 2,
		part: 1,
		inputFile: "inputs/day_02_input.txt",
		fn: func(input string) any {return day_02.PartOne(input)},
	},
	{
		day: 2,
		part: 2,
		inputFile: "inputs/day_02_input.txt",
		fn: func(input string) any {return day_02.PartTwo(input)},
	},
	{
		day: 3,
		part: 1,
		inputFile: "inputs/day_03_input.txt",
		fn: func(input string) any {return day_03.PartOne(input)},
	},
	{
		day: 3,
		part: 2,
		inputFile: "inputs/day_03_input.txt",
		fn: func(input string) any {return day_03.PartTwo(input)},
	},
	{
		day: 4,
		part: 1,
		inputFile: "inputs/day_04_input.txt",
		fn: func(input string) any {return day_04.PartOne(input)},
	},
	{
		day: 4,
		part: 2,
		inputFile: "inputs/day_04_input.txt",
		fn: func(input string) any {return day_04.PartTwo(input)},
	},
	{
		day: 5,
		part: 1,
		inputFile: "inputs/day_05_input.txt",
		fn: func(input string) any {return day_05.PartOne(input)},
	},
	{
		day: 5,
		part: 2,
		inputFile: "inputs/day_05_input.txt",
		fn: func(input string) any {return day_05.PartTwo(input)},
	},
}

type runFunc func(string) any

type puzzleData struct {
	day int
	part int
	inputFile string
	fn runFunc
}

func getFileContents(filePath string) string {
	fileBuf, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(fileBuf)
}

type puzzleResult struct {
	day int
	part int
	output string
}

func runAll(puzzles []puzzleData) []puzzleResult {
	results := make([]puzzleResult, 0, len(puzzles))

	for _, d := range puzzles {
		input := getFileContents(d.inputFile)
		results = append(
			results,
			puzzleResult{
				day: d.day,
				part: d.part,
				output: fmt.Sprint(d.fn(input)),
			},
		)
	}

	return results
}
