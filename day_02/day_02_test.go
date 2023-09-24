package day_02

import (
	"testing"
	"local-advent-2022/util"
)

var TEST_INPUT string =
`A Y
B X
C Z`

func TestPart1TestInput(t *testing.T) {
	result := PartOne(TEST_INPUT)
	expectedScore := uint(15)

	if expectedScore != result {
		t.Fatalf("Incorrect scoring. Expected %d, got %d", expectedScore, result)
	}
}

func TestPart2TestInput(t *testing.T) {
	result := PartTwo(TEST_INPUT)
	expectedScore := uint(12)

	if expectedScore != result {
		t.Fatalf("Incorrect outcome scoring. Expected %d, got %d", expectedScore, result)
	}
}

func TestDecodedPlays(t *testing.T) {
	lines := util.AsLines(TEST_INPUT)
	plays := decodeAllPlays(lines)

	expectedPlays := []rpsPlay{
		{
			opponentMove: ROCK,
			myMove: PAPER,
		},
		{
			opponentMove: PAPER,
			myMove: ROCK,
		},
		{
			opponentMove: SCISSORS,
			myMove: SCISSORS,
		},
	}
	if len(expectedPlays) != len(plays) {
		t.Errorf(
			"Mismatched different number of parsed plays: expected %d, got %d",
			len(expectedPlays),
			len(plays),
		)
	}

	for i, expectedPlay := range expectedPlays {
		actualPlay := plays[i]
		mismatched :=
			expectedPlay.opponentMove != actualPlay.opponentMove ||
			expectedPlay.myMove != actualPlay.myMove
		if mismatched {
			t.Errorf("Mismatched plays: expected %v, got %v", expectedPlay, actualPlay)
		}
	}

	if t.Failed() {
		printExpectedMoves(expectedPlays, plays, t)
		t.FailNow()
	}
}

func printExpectedMoves(expected []rpsPlay, actual []rpsPlay, t *testing.T) {
	t.Logf("Expected: %v\nGot: %v\n", expected, actual)
}
