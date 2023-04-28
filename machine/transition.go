package machine

const Right = iota
const Left = iota

type Transition struct {
	read      rune
	write     rune
	direction int
	endState  *State
}
