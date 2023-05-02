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

	machine := tu.Machine{Tape: &tape}
	instructions := strings.Split(machineString, "11")

	for _, v := range instructions {
		err := processMachineInstruction(&machine, v)
		if err != nil {
			return machine, err
		}
	}

	setStartElement(&machine)
	return machine, nil
}

func processMachineInstruction(machine *tu.Machine, instruction string) error {
	counters := strings.Split(instruction, "1")

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
	re := regexp.MustCompile("^0+$")

	for i := 0; i < len(inst); i++ {
		if inst[i] == "" {
			remove(inst, i)
		}
	}

	for _, v := range inst {
		if !re.MatchString(v) {
			return errors.New("string does not match pattern")
		}
	}
	return nil
}

func readDirection(i int) (int, error) {
	if i == 2 {
		return tu.Right, nil
	} else if i == 1 {
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

func setStartElement(machine *tu.Machine) {
	for _, v := range machine.States {
		if v.Number == 1 {
			machine.CurrentState = v
		}
	}

}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
