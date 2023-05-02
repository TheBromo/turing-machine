package main

import (
	"fmt"
	"regexp"
	"time"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

//TODO binary number representation
//TODO step and laufmodus

func main() {
	//read file
	input := "010010001010011000101010010110001001001010011000100010001010111100011111"

	re := regexp.MustCompile("111")

	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}
		for {
			logger.PrintMachine(turing, err)
			err = turing.DoStep()
			time.Sleep(2 * time.Second) // pauses execution for 2 seconds
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
