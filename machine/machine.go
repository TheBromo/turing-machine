package machine

import "strconv"

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

func (machine *Machine) ToString() string {
	return machine.tapeToString() + "\n" + machine.stateToString()
}

func (machine *Machine) stateToString() string {
	state := ""
	current := strconv.Itoa(machine.CurrentState.Number)
	state += "q" + current
	for _, v := range machine.CurrentState.Transitions {
		read := strconv.Itoa(v.Read)
		end := strconv.Itoa(v.EndState.Number)
		write := strconv.Itoa(v.Write)
		dir := ""
		if v.Direction == Left {
			dir += "L"
		} else {
			dir += "R"
		}
		state += "\n" + "δ(q" + current + ", " + read + ") = (q" + end + "," + write + ", " + dir + "),"
	}
	return state
}

func (machine *Machine) tapeToString() string {

	tape := ""
	arrow := ""

	for i, v := range machine.Tape.Content {
		tape += " " + strconv.Itoa(v)
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
