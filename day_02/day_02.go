package day_02

import (
	"fmt"
	"strings"
	"local-advent-2022/util"
)

func PartOne(input string) uint {
	lines := util.AsLines(input)
	plays := decodeAllPlays(lines)
	return scorePlays(plays)
}

func PartTwo(input string) uint {
	lines := util.AsLines(input)
	outcomes := decodeAllOutcomes(lines)
	return scoreOutcomes(outcomes)
}

type rpsMove uint
const (
	ROCK rpsMove = iota
	PAPER
	SCISSORS
)
type rpsMoveInfo struct{
	sign rpsMove
	value uint
	beats rpsMove
	losesTo rpsMove
}

type rpsPlay struct {
	opponentMove rpsMove
	myMove rpsMove
}

type gameOutcome uint
const (
	WIN gameOutcome = iota
	DRAW
	LOSS
)
type rpsWantedOutcome struct {
	opponentMove rpsMove
	desiredOutcome gameOutcome
}

type rpsMoveData struct{
	moves map[rpsMove]rpsMoveInfo
}

func newRpsMoveData() rpsMoveData{
	return rpsMoveData{
		moves: map[rpsMove]rpsMoveInfo{
			ROCK: {
				sign: ROCK,
				value: 1,
				beats: SCISSORS,
				losesTo: PAPER,
			},
			PAPER: {
				sign: PAPER,
				value: 2,
				beats: ROCK,
				losesTo: SCISSORS,
			},
			SCISSORS: {
				sign: SCISSORS,
				value: 3,
				beats: PAPER,
				losesTo: ROCK,
			},
		},
	}
}

func (moveData *rpsMoveData) scorePlay(play rpsPlay) uint {
	myMove := moveData.moves[play.myMove]

	if myMove.beats == play.opponentMove {
		return 6 + myMove.value // win
	} else if myMove.sign == play.opponentMove {
		return 3 + myMove.value // draw
	} else if myMove.losesTo == play.opponentMove {
		return myMove.value // loss
	} else {
		panic(fmt.Sprintf("Illegal state, outcome of play undetermined: %v", play))
	}
}

func (moveData *rpsMoveData) scoreOutcome(outcome rpsWantedOutcome) uint {
	opponentMove := moveData.moves[outcome.opponentMove]

	var myMove rpsMoveInfo
	score := uint(0)
	switch outcome.desiredOutcome {
	case WIN:
		myMove = moveData.moves[opponentMove.losesTo]
		score = 6
	case DRAW:
		myMove = opponentMove
		score = 3
	case LOSS:
		myMove = moveData.moves[opponentMove.beats]
		score = 0
	default:
		panic(fmt.Sprintf("Illegal state, play for outcome undetermined: %v", outcome))
	}

	score += myMove.value
	return score
}

func mapOpponentMove(o rune) rpsMove {
	switch o {
	case 'A':
		return ROCK
	case 'B':
		return PAPER
	case 'C':
		return SCISSORS
	default:
		panic(fmt.Sprintf("got an unknown character representing an opponent's move: %s", string(o)))
	}
}

func mapMyMove(o rune) rpsMove {
	switch o {
	case 'X':
		return ROCK
	case 'Y':
		return PAPER
	case 'Z':
		return SCISSORS
	default:
		panic(fmt.Sprintf("got an unknown character representing my move: %s", string(o)))
	}
}

func mapDesiredOutcome(o rune) gameOutcome {
	switch o {
	case 'X':
		return LOSS
	case 'Y':
		return DRAW
	case 'Z':
		return WIN
	default:
		panic(fmt.Sprintf("got an unknown character representing an outcome: %s", string(o)))
	}
}

func decodeMoves(input string) rpsPlay {
	encodedMoves := strings.Fields(input)
	if len(encodedMoves) != 2 {
		panic(fmt.Sprintf(
			"Encountered a play encoding that did not have exactly two items: \"%s\"",
			input,
		))
	}

	oppEncMove := []rune(encodedMoves[0])
	myEncMove := []rune(encodedMoves[1])
	if len(oppEncMove) != 1 {
		panic(fmt.Sprintf("Encoded opponent move not valid: \"%s\"", string(oppEncMove)))
	}
	if len(myEncMove) != 1 {
		panic(fmt.Sprintf("Encoded my move not valid: \"%s\"", string(myEncMove)))
	}

	return rpsPlay{
		opponentMove: mapOpponentMove(oppEncMove[0]),
		myMove: mapMyMove(myEncMove[0]),
	}
}

func decodeAllPlays(lines []string) []rpsPlay {
	decodedPlays := make([]rpsPlay, 0, len(lines))
	for _, line := range lines {
		decodedPlays = append(decodedPlays, decodeMoves(line))
	}
	return decodedPlays
}

func scorePlays(plays []rpsPlay) uint {
	moveData := newRpsMoveData()
	total := uint(0)
	for _, play := range plays {
		total += moveData.scorePlay(play)
	}
	return total
}

func decodeOutcomes(input string) rpsWantedOutcome {
	encodedRunes := strings.Fields(input)
	if len(encodedRunes) != 2 {
		panic(fmt.Sprintf(
			"Encountered a play encoding that did not have exactly two items: \"%s\"",
			input,
		))
	}

	oppEncMove := []rune(encodedRunes[0])
	encOutcome := []rune(encodedRunes[1])
	if len(oppEncMove) != 1 {
		panic(fmt.Sprintf("Encoded opponent move not valid: \"%s\"", string(oppEncMove)))
	}
	if len(encOutcome) != 1 {
		panic(fmt.Sprintf("Encoded outcome not valid: \"%s\"", string(encOutcome)))
	}

	return rpsWantedOutcome{
		opponentMove: mapOpponentMove(oppEncMove[0]),
		desiredOutcome: mapDesiredOutcome(encOutcome[0]),
	}
}

func decodeAllOutcomes(lines []string) []rpsWantedOutcome {
	decodedOutcomes := make([]rpsWantedOutcome, 0, len(lines))
	for _, line := range lines {
		decodedOutcomes = append(decodedOutcomes, decodeOutcomes(line))
	}
	return decodedOutcomes
}

func scoreOutcomes(outcomes []rpsWantedOutcome) uint {
	moveData := newRpsMoveData()
	total := uint(0)
	for _, outcome := range outcomes {
		total += moveData.scoreOutcome(outcome)
	}
	return total
}
