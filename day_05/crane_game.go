package day_05

type craneGame struct {
	stacks []stack
}

type craneMove struct {
	src uint
	dest uint
	count uint
}

func makeWithColumns(numCols int) *craneGame {
	gameState := &craneGame{
		stacks: make([]stack, numCols),
	}

	return gameState
}

// move a value/crate from src to dest
// src and dest are 1-indexed
func (cg *craneGame) move(mv craneMove) {
	for i := uint(0); i < mv.count; i++ {
		val, success := cg.stacks[mv.src - 1].pop()
		if !success {
			return
		}

		cg.stacks[mv.dest - 1].push(val)
	}
}

// instead of moving one crate at a time, moves a block without changing their order
func (cg *craneGame) moveSubStack(mv craneMove) {
	block := make([]byte, 0, mv.count)
	for i := uint(0); i < mv.count; i++ {
		val, success := cg.stacks[mv.src -1].pop()
		if !success {
			return
		}

		block = append(block, val)
	}

	for i := len(block) - 1; i >= 0; i-- {
		cg.stacks[mv.dest -1].push(block[i])
	}
}

type stack []byte

func emptyStack() stack {
	return make([]byte, 0)
}

func (s *stack) push(v byte) {
	*s = append(*s, v)
}

func (s *stack) pop() (val byte, success bool) {
	val, success = s.peek()

	if success {
		*s = (*s)[:len(*s) - 1]
	}
	return val, true
}
func (s *stack) peek() (val byte, success bool) {
	size := len(*s)
	if size == 0 {
		return 0, false
	}

	val = (*s)[size -1]
	return val, true
}
