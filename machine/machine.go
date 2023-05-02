package machine

type Machine struct {
	CurrentState *State
	States       []*State
	Tape         *Tape
}

func (machine *Machine) DoStep() error {
	input := machine.Tape.Read()
	transition, err := machine.CurrentState.getTransitionForInput(input)
	if err != nil {
		return err
	}
	machine.Tape.Move(transition.Direction, transition.Write)
	machine.CurrentState = transition.EndState
	return err
}

func (machine *Machine) GetOrAddState(state State) *State {
	for _, e := range machine.States {
		if e.Number == state.Number {
			return e
		}
	}
	machine.addNewState(&state)
	return &state
}

func (machine Machine) hasState(state State) bool {
	for _, e := range machine.States {
		if e.Number == state.Number {
			return true
		}
	}
	return false
}

func (machine *Machine) addNewState(state *State) {
	if !machine.hasState(*state) {
		machine.States = append(machine.States, state)
	}
}
