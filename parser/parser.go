package machine

import (
	"errors"
	"regexp"

	tu "github.com/thebromo/turing-machine/machine"
)

func InitTape(tape string) tu.Tape {
	t := tu.Tape{}
	t.Content = []rune(tape)
	t.Position = 0
	return t
}



func InitMachine(machineString string) (tu.Machine, error) {
	atEnd := false
	re := regexp.MustCompile(`(0+1){5}`)
	machine := tu.Machine{}

	for !atEnd {
		loc := re.FindIndex([]byte(machineString))

		if len(loc) == 0 {
			atEnd = true
		}

		instruction := machineString[loc[0]:loc[1]]
		machineString = machineString[loc[1]:]

		err := processMachineInstruction(&machine, instruction)

		if err != nil {
			return machine, err
		}

	}

	return machine, nil
}

func processMachineInstruction(machine *tu.Machine, instruction string) error {
	counters := regexp.MustCompile("1").Split(instruction, 5)

	if len(counters) != 5 {
		return errors.New("incorrect instruction")
	}

	startState := machine.GetOrAddState(readState(len(counters[0])))
	read := readBinaryOperator(len(counters[1]))
	endState := machine.GetOrAddState(readState(len(counters[2])))
	write := readBinaryOperator(len(counters[3]))
	direction := readDirection(len(counters[4]))

	transit := tu.Transition{
		Read:      rune(read),
		Write:     rune(write),
		EndState:  endState,
		Direction: direction,
	}
	startState.Transitions = append(startState.Transitions, transit)

	return nil
}

func readDirection(i int) int {
	if i == 0 {
		return tu.Right
	} else {
		return tu.Left
	}
}

func readState(id int) tu.State {
	state := tu.State{Number: id, Transitions: make([]tu.Transition, 0)}
	return state
}

func readBinaryOperator(count int) int {
	return count - 1
}
