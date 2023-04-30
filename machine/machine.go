package machine

type Machine struct {
	CurrentState *State
	States       []*State
	Tape         *Tape
}

func (machine *Machine) DoStep() error {
	//TODO set current state
	input := machine.Tape.Read()
	transition, err := machine.CurrentState.getTransitionForInput(input)
	if err != nil {
		return err
	}
	machine.Tape.Move(transition.Direction, transition.Write)
	machine.CurrentState = transition.EndState
	return nil
}

func (machine *Machine) TapeToString() string {
	tape := ""
	arrow := ""

	for i, v := range machine.Tape.Content {
		tape += " " + string(v)
		if i == machine.Tape.Position {
			arrow += " ▲"
		} else {
			arrow += "  "
		}

	}
	return tape + "\n" + arrow
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
