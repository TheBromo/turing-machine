package machine

const fillerSymbol int = 5

type Tape struct {
	Content  []int
	Position int
}

func (t *Tape) Read() int {
	if t.Position < 0 {
		t.Content = append(append(make([]int, 1), fillerSymbol), t.Content...)
		t.Position = 0
	} else if t.Position >= len(t.Content) {
		t.Content = append(t.Content, fillerSymbol)
	}
	return t.Content[t.Position]
}

func (t *Tape) Move(direction int, write int) {

	t.Content[t.Position] = write

	if direction == Right {
		t.Position += 1
	} else {
		t.Position -= 1
	}
}
