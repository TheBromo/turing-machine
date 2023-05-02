package machine

type Tape struct {
	Content  []int
	Position int
}

func (t *Tape) Read() int {
	if t.Position < 0 || t.Position > len(t.Content) {
		return -1
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
