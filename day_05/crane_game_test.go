package day_05

import (
	"testing"
)

func TestMoves(t *testing.T) {
	testGame := craneGame{
		stacks: []stack{
			{'a', 'b', 'c'},
			{},
		},
	}

	testGame.move(craneMove{src: 1, dest: 2, count: 1})
	if 1 != len(testGame.stacks[1]) {
		t.Fatalf(
			"Expected the second stack to have a size of 1, but stack was %s",
			string(testGame.stacks[1]),
		)
	}
	if testGame.stacks[1][0] != 'c' {
		t.Fatalf(
			"Expected there to be a 'c' value in the second stack, but stack was %s",
			string(testGame.stacks[1]),
		)
	}

	testGame.move(craneMove{src: 1, dest: 2, count: 2})

	if 3 != len(testGame.stacks[1]) {
		t.Fatalf(
			"Expected the second stack to have a size of 3, but stack was %s",
			string(testGame.stacks[1]),
		)
	}

	if testGame.stacks[1][1] != 'b' {
		t.Fatalf(
			"Expected there to be a 'b' value in the second stack, but stack was %s",
			string(testGame.stacks[1]),
		)
	}
	if testGame.stacks[1][2] != 'a' {
		t.Fatalf(
			"Expected there to be a 'a' value in the second stack, but stack was %s",
			string(testGame.stacks[1]),
		)
	}
}
