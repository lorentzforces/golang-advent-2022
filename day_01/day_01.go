package day_01

import (
	"fmt"
	"sort"
	"strconv"
	"local-advent-2022/util"
)

func PartOne(input string) uint {
	lines := util.AsLines(input)
	elfValues := calcElfValues(lines)
	return getMaxValue(elfValues)
}

func PartTwo(input string) uint {
	lines := util.AsLines(input)
	elfValues := calcElfValues(lines)
	sort.Slice(
		elfValues,
		func (i, j int) bool {return !(elfValues[i] < elfValues[j])}, // reverse order
	)

	topThree := elfValues[:3]
	return topThree[0] + topThree[1] + topThree[2]
}

func calcElfValues(lines []string) []uint {
	elfValues := make([]uint, 1)

	for _, line := range lines {
		if len(line) > 0 {
			val, err := strconv.ParseUint(line, 10, 0)
			if (err != nil) {
				panic(fmt.Sprintf(
					"Encountered non-blank line which was not an unsigned integer: %s\n",
					line,
				))
			}

			elfValues[len(elfValues) - 1] += uint(val)
		} else {
			if elfValues[len(elfValues) - 1] != 1 {
				elfValues = append(elfValues, 0)
			}
		}
	}

	return elfValues
}

func getMaxValue(vals []uint) uint {
	var maxVal uint = 0;

	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}
