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
	input := "01001000101001100010101001011000100100101001100010001000101011100110101010"

	re := regexp.MustCompile("111")

	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}

		for turing.DoStep() != nil {
			logger.Print(turing.TapeToString())
			time.Sleep(2 * time.Second) // pauses execution for 2 seconds
		}

		fmt.Printf("turing: %v\n", turing)
		fmt.Printf("tape: %v\n", tape)

	} else {
		fmt.Errorf("incorrect format")
	}

}
