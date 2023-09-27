package day_03

import "testing"

var TEST_INPUT string =
`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1TestInput(t *testing.T) {
	result := PartOne(TEST_INPUT)
	expectedResult := uint(157)

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %d, got %d", expectedResult, result)
	}
}

func TestPart2TestInput(t *testing.T) {
	result := PartTwo(TEST_INPUT)
	expectedResult := uint(70)

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %d, got %d", expectedResult, result)
	}
}

func TestPriorityValues(t *testing.T) {
	testPriorityValue('a', 1, t)
	testPriorityValue('z', 26, t)
	testPriorityValue('A', 27, t)
	testPriorityValue('Z', 52, t)
}

func testPriorityValue(char byte, expected byte, t *testing.T) {
	val, err := itemPriority(char)
	if err != nil {
		t.Errorf("Error encountered getting priority for '%c': %v", char, err)
	}
	if expected != val {
		t.Errorf(
			"itemPriority produced unexpected value for '%c'. Expected %d, got %d",
			char,
			expected,
			val,
		)
	}
}
