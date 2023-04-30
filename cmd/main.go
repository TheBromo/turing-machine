package main

import logger "github.com/thebromo/turing-machine/log"

// import (
// 	"fmt"
// 	"regexp"

// 	parser "github.com/thebromo/turing-machine/parser"
// )

func main() {
	logger.Print()
	//read file
	//input := "0010101001010100101011101010101"

	//re := regexp.MustCompile(".+111.+")

	//split := re.Split(input, 0)

	// if len(split) == 2 {
	// 	turing, err := parser.InitMachine(split[0])
	// 	tape := parser.InitTape(split[1])

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("turing: %v\n", turing)
	// 	fmt.Printf("tape: %v\n", tape)

	// } else {
	// 	fmt.Errorf("incorrect format")
	// }

}
