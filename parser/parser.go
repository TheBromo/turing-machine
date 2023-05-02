package machine

import (
	"errors"
	"regexp"
	"strings"

	tu "github.com/thebromo/turing-machine/machine"
)

func InitTape(tape string) tu.Tape {
	t := tu.Tape{}

	t.Content = make([]int, len(tape))

	for i, v := range tape {
		t.Content[i] = int(v - '0')
	}

	t.Position = 0
	return t
}

func InitMachine(machineString string, tape tu.Tape) (tu.Machine, error) {

	re := regexp.MustCompile(`(0+1){5}`)
	machine := tu.Machine{Tape: &tape}

	for len(machineString) > 0 {
		loc := re.FindIndex([]byte(machineString))

		if len(loc) == 0 {
			machine.CurrentState = machine.States[0]
			return machine, nil
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

	for i, v := range counters {
		counters[i] = strings.ReplaceAll(v, "1", "")
	}

	//add erors
	err := checkForIncorrectInstructions(counters)

	if err != nil {
		return err
	}

	startState := machine.GetOrAddState(readState(len(counters[0])))
	read := readBinaryOperator(len(counters[1]))
	endState := machine.GetOrAddState(readState(len(counters[2])))
	write := readBinaryOperator(len(counters[3]))
	direction, err := readDirection(len(counters[4]))

	if err != nil {
		return err
	}

	transit := tu.Transition{
		Read:      read,
		Write:     write,
		EndState:  endState,
		Direction: direction,
	}
	startState.Transitions = append(startState.Transitions, transit)

	return nil
}

func checkForIncorrectInstructions(inst []string) error {
	re := regexp.MustCompile("^(1|0)+$")

	for _, v := range inst {
		if !re.MatchString(v) {
			return errors.New("string does not match pattern")
		}
	}
	return nil
}

func readDirection(i int) (int, error) {
	if i == 1 {
		return tu.Right, nil
	} else if i == 2 {
		return tu.Left, nil
	}
	return -1, errors.New("incorrect direction")
}

func readState(id int) tu.State {
	state := tu.State{Number: id, Transitions: make([]tu.Transition, 0)}
	return state
}

func readBinaryOperator(count int) int {
	return count - 1
}
