package machine

const (
    Left = iota
    Right 
)

type Transition struct {
	Read      rune
	Write     rune
	Direction int
	EndState  *State
}
