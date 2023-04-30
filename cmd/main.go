package main

import (
	"fmt"
	"regexp"

	machine "github.com/thebromo/turing-machine/machine"
)

func main() {
	//read file
	input := "0010101001010100101011101010101"

	re := regexp.MustCompile(".+111.+")

	split := re.Split(input, 0)

	if len(split) == 2 {
		turing, err := machine.InitMachine(split[0])
		tape := machine.InitTape(split[1])

		if err != nil {
			panic(err)
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
