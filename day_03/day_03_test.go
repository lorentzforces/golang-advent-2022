package day_03

import "testing"

var TEST_INPUT string =
`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestCreateInventoryCreatesExpectedValues(t *testing.T) {
	result, err := parseInventory("abcabd")
	if err != nil {
		t.Fatalf("parseInventory produced error: %v", err)
	}

	expectedLeft, expectedRight := "abc", "abd"
	expectedBoth := "ab"
	if expectedLeft != result.left {
		t.Errorf("Left side unexpected value. Expected %s, got %s", expectedLeft, result.left)
	}
	if expectedRight != result.right {
		t.Errorf("Right side unexpected value. Expected %s, got %s", expectedRight, result.right)
	}

	if expectedBoth != result.both {
		t.Errorf("Both side unexpected value. Expected %s, got %s", expectedBoth, result.both)
	}

	sortedResult, err := parseInventory("cbadab")
	if err != nil {
		t.Fatalf("parseInventory produced error: %v", err)
	}
	if expectedLeft != sortedResult.left {
		t.Errorf(
			"Left side unexpected sorted value. Expected %s, got %s",
			expectedLeft,
			sortedResult.left,
		)
	}
	if expectedRight != sortedResult.right {
		t.Errorf(
			"Right side unexpected sorted value. Expected %s, got %s",
			expectedRight,
			sortedResult.right,
		)
	}
	if expectedBoth != sortedResult.both {
		t.Errorf(
			"Both side unexpected sorted value. Expected %s, got %s",
			expectedBoth,
			sortedResult.both,
		)
	}

	dedupedLeft, dedupedRight, dedupedBoth := "bc", "bc", "bc"
	dedupedResult, err := parseInventory("ccbbccbb")
	if err != nil {
		t.Fatalf("parseInventory produced error: %v", err)
	}
	if dedupedLeft != dedupedResult.left {
		t.Errorf(
			"Left side unexpected deduped value. Expected %s, got %s",
			dedupedLeft,
			dedupedResult.left,
		)
	}
	if dedupedRight != dedupedResult.right {
		t.Errorf(
			"Right side unexpected deduped value. Expected %s, got %s",
			dedupedRight,
			dedupedResult.right,
		)
	}
	if dedupedBoth != dedupedResult.both {
		t.Errorf(
			"Both side unexpected deduped value. Expected %s, got %s",
			dedupedBoth,
			dedupedResult.both,
		)
	}

}

func TestPriorityValues(t *testing.T) {
	testPriorityValue('a', 1, t)
	testPriorityValue('z', 26, t)
	testPriorityValue('A', 27, t)
	testPriorityValue('Z', 52, t)
}

func testPriorityValue(char byte, expected uint, t *testing.T) {
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

func TestParat1TestInput(t *testing.T) {
	result := PartOne(TEST_INPUT)
	expectedResult := uint(157)

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %d, got %d", expectedResult, result)
	}
}
