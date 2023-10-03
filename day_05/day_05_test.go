package day_05

import (
	"testing"
)

var TEST_INPUT string =
`    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1TestInput(t *testing.T) {
	result := PartOne(TEST_INPUT)
	expectedResult := "CMZ"

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %s, got %s\n", expectedResult, result)
	}
}

func TestPart2TestInput(t *testing.T) {
	result := PartTwo(TEST_INPUT)
	expectedResult := "MCD"

	if expectedResult != result {
		t.Fatalf("Incorrect value produced. Expected %s, got %s\n", expectedResult, result)
	}
}

func Test_ParseMove_HappyPath(t *testing.T) {
	line := "move 1 from 2 to 3"
	move, err := parseMove(line)
	if err != nil {
		t.Fatal(err)
	}

	expectedMove := craneMove{
		count: 1,
		src: 2,
		dest: 3,
	}

	if expectedMove.count != move.count {
		t.Errorf("Move count should be %d, but got: %d", expectedMove.count, move.count)
	}
	if expectedMove.src != move.src {
		t.Errorf("Move src should be %d, but got: %d", expectedMove.src, move.src)
	}
	if expectedMove.dest != move.dest {
		t.Errorf("Move dest should be %d, but got: %d", expectedMove.dest, move.dest)
	}
}
