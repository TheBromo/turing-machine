package machine

import (
	"fmt"
)

type Machine struct {
	states []State
	tape   Tape
}

func (m *Machine) Load(machine string) error {
	tempCount := 0

	for i := 0; i < len(machine); i++ {
		if machine[i] == '0' {

		} else if machine[i] == '1' {

		} else {
			return fmt.Errorf("Incorrect Symbol %s", machine[i])
		}
	}

	return nil
}
