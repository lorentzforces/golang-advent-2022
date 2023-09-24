package day_01

import (
	"testing"
	"local-advent-2022/util"
)

var TEST_INPUT string =
`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPartOneInput(t *testing.T) {
	result := PartOne(TEST_INPUT)

	if result != 24000 {
		t.Fatalf("Day 1 part 1 test input failed. Expected %d, got %d", 24000, result)
	}
}

func TestPartTwoInput(t *testing.T) {
	result := PartTwo(TEST_INPUT)

	if result != 45000 {
		t.Fatalf("Day 1 part 2 test input failed. Expected %d, got %d", 45000, result)
	}
}

func TestCalcSampleValues(t *testing.T) {
	lines := util.AsLines(TEST_INPUT)
	calculatedValues := calcElfValues(lines)

	expectedValues := []uint {
		6000,
		4000,
		11000,
		24000,
		10000,
	}

	assertSlicesEqual(expectedValues, calculatedValues, t)
}

func assertSlicesEqual(expected []uint, actual []uint, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Slice lengths differed. Expected %d, got %d", len(expected), len(actual))
		printExpectedSlices(expected, actual, t)
		t.FailNow()
	}

	for i, expectedVal := range expected {
		if expectedVal != actual[i] {
			t.Errorf(
				"Slice elements differed at index %d. Expected %d, got %d",
				i,
				expectedVal,
				actual[i],
			)
		}
	}
	if t.Failed() {
		printExpectedSlices(expected, actual, t)
		t.FailNow()
	}
}

func printExpectedSlices(expected []uint, actual []uint, t *testing.T) {
	t.Logf("Expected: %v\nGot: %v\n", expected, actual)
}
