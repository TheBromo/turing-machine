package machine

import (
	"errors"
)

type State struct {
	Number      int
	Transitions []Transition
}

func (state *State) getTransitionForInput(input rune) (Transition, error) {
	for _, v := range state.Transitions {
		if v.Read == input {
			return v, nil
		}
	}
	return Transition{}, errors.New("no transition found for input")
}
