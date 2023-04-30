package machine

const Right = iota
const Left = iota

type Transition struct {
	Read      rune
	Write     rune
	Direction int
	EndState  *State
}
