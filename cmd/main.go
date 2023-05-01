package main

import (
	"fmt"
	"regexp"
	"time"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

func main() {
	//read file
	input := "101001010010011010100010100110001001010010011000101001010011110110101010"

	re := regexp.MustCompile("111")

	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}
		logger.Print(turing.ToString())
		for turing.DoStep() == nil {
			logger.Print(turing.ToString())
			time.Sleep(2 * time.Second) // pauses execution for 2 seconds
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
