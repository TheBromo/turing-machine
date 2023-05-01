package machine

const (
    Left = iota
    Right 
)

type Transition struct {
	Read      int
	Write     int
	Direction int
	EndState  *State
}
