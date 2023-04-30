package machine

import (
	"errors"
)

type Tape struct {
	Content  []int
	Position int
}

func (t *Tape) Read() int {
	return t.Content[t.Position]
}

func (t *Tape) Move(direction int, write int) (int, error) {
	if t.Position >= 0 && t.Position < len(t.Content) {
		t.Content[t.Position] = write

		if direction == Right {
			t.Position = t.Position - 1
		} else {
			t.Position = t.Position + 1
		}
		return t.Read(), nil
	} else {
		return 0, errors.New("position out of bounds")
	}
}
