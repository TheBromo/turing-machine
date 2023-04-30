package machine

import (
	"errors"
	"regexp"
)

type Machine struct {
	states []*State
	tape   *Tape
}

//Regex ^(0|1){3,}111(0|1){3,}$
//instruction: (0+1){3}
//TODO maybe instead of three its 4

func InitMachine(machineString string) (Machine, error) {
	atEnd := false
	re := regexp.MustCompile(`(0+1){5}`)
	machine := Machine{}

	for !atEnd {
		loc := re.FindIndex([]byte(machineString))

		if len(loc) == 0 {
			atEnd = true
		}

		instruction := machineString[loc[0]:loc[1]]
		machineString = machineString[loc[1]:]

		err := processInstruction(&machine, instruction)

		if err != nil {
			return machine, err
		}

	}

	return machine, nil
}

func processInstruction(machine *Machine, instruction string) error {
	counters := regexp.MustCompile("1").Split(instruction, 5)

	if len(counters) != 5 {
		return errors.New("incorrect instruction")
	}

	startState := machine.getOrAddState(readState(len(counters[0])))
	read := readBinaryOperator(len(counters[1]))
	endState := machine.getOrAddState(readState(len(counters[2])))
	write := readBinaryOperator(len(counters[3]))
	direction := readDirection(len(counters[4]))

	transit := Transition{
		read:      rune(read),
		write:     rune(write),
		endState:  endState,
		direction: direction,
	}
	startState.transitions = append(startState.transitions, transit)

	return nil
}

func (machine Machine) hasState(state State) bool {
	for _, e := range machine.states {
		if e.number == state.number {
			return true
		}
	}
	return false
}

func (machine *Machine) getOrAddState(state State) *State {
	for _, e := range machine.states {
		if e.number == state.number {
			return e
		}
	}
	machine.addNewState(&state)
	return &state
}

func (machine *Machine) addNewState(state *State) {
	if !machine.hasState(*state) {
		machine.states = append(machine.states, state)
	}
}

func readDirection(i int) int {
	if i == 0 {
		return Right
	} else {
		return Left
	}
}

func readState(id int) State {
	state := State{number: id, transitions: make([]Transition, 0)}
	return state
}

func readBinaryOperator(count int) int {
	return count - 1
}
