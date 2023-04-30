package main

import (
	"fmt"
	"regexp"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

func main() {
	//read file
	input := "1010010001010011000101010010110001001001010011000100010001010111110100100101010101"

	re := regexp.MustCompile("111")

	split := re.Split(input, 2)

	if len(split) == 2 {
		turing, err := parser.InitMachine(split[0])
		tape := parser.InitTape(split[1])

		if err != nil {
			panic(err)
		}

		for turing.DoStep() != nil {
			logger.Print(turing.TapeToString())
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
