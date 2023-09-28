package day_04

import (
	"fmt"
	"local-advent-2022/util"
	"testing"
)

var TEST_INPUT string =
`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1TestInput(t *testing.T) {
	result := PartOne(TEST_INPUT)
	expectedResult := uint(2)

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %d, got %d", expectedResult, result)
	}
}

func TestPart2TestInput(t *testing.T) {
	result := PartTwo(TEST_INPUT)
	expectedResult := uint(4)

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %d, got %d", expectedResult, result)
	}
}

// not a functional test, debugging only with -v
func TestSpanPairParsing(t *testing.T) {
	lines := util.AsLines(TEST_INPUT)
	for _, line := range lines {
		result, err := parseSpanPair(line)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("Parsed line: \"%s\"\n  Result: %v\n", line, result)
	}
}
