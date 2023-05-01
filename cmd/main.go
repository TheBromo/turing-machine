package main

import (
	"fmt"
	"regexp"
	"time"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

//TODO left and right
//TODO check if states are read correctly

func main() {
	//read file
	input := "01001000101001 10001010100101 1000100100101001 1000100010001010 11110110101010"

	re := regexp.MustCompile("111")

	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}
		logger.PrintMachine(turing)
		for turing.DoStep() == nil {
			logger.PrintMachine(turing)
			time.Sleep(2 * time.Second) // pauses execution for 2 seconds
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
