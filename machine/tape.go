package machine

import (
	"errors"
)

type Tape struct {
	Content  []rune
	Position int
}

func (t *Tape) Read() rune {
	return t.Content[t.Position]
}

func (t *Tape) Move(direction int, write rune) (rune, error) {
	if t.Position >= 0 && t.Position < len(t.Content) {
		t.Content[t.Position] = write

		if direction == Right {
			t.Position = t.Position - 1
		} else {
			t.Position = t.Position + 1
		}
		return t.Read(), nil
	} else {
		return rune(0), errors.New("position out of bounds")
	}
}

func (t *Tape) Init(tape string) {
    t.Content = []rune(tape);
    t.Position = 0
}
