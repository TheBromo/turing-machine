package machine

type Machine struct {
	currentState *State
	states       []*State
	tape         *Tape
}

func (machine *Machine) DoStep() error {
	input := machine.tape.Read()
	transition, err := machine.currentState.getTransitionForInput(input)
	if err != nil {
		return err
	}
	machine.tape.Move(transition.Direction, transition.Write)
	machine.currentState = transition.EndState
	return nil
}

func (machine *Machine) ToString() {
	//visualize tape
	//visualize states

    //
}

func (machine *Machine) GetOrAddState(state State) *State {
	for _, e := range machine.states {
		if e.Number == state.Number {
			return e
		}
	}
	machine.addNewState(&state)
	return &state
}

func (machine Machine) hasState(state State) bool {
	for _, e := range machine.states {
		if e.Number == state.Number {
			return true
		}
	}
	return false
}

func (machine *Machine) addNewState(state *State) {
	if !machine.hasState(*state) {
		machine.states = append(machine.states, state)
	}
}
