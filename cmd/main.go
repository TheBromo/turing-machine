package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

//TODO binary number representation
//TODO step and laufmodus

func main() {
	fmt.Print("Stepmpdus (y/N)?\n> ")
	stepModus := askBool()
	file := askFile()
	fmt.Print(file)
	//read file
	input := "0100100010100110001010100101100010010010100110001000100010101110000010000"
	re := regexp.MustCompile("111")
	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}
		for err == nil {
			logger.PrintMachine(turing, err)

			err = turing.DoStep()
			if stepModus {
				time.Sleep(2 * time.Second) // pauses execution for 2 seconds
			}
		}

		logger.PrintMachine(turing, err)

	} else {
		fmt.Errorf("incorrect format")
	}
}

func askBool() bool {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'y', 'Y':
		return true
	default:
		return false
	}

}

func askFile() string {
	path := "C:\\Users\\manue\\IdeaProjects\\turing-machine\\machine.turing"
	fmt.Println("enter custom file (y/N)")
	ask := askBool()

	if ask {
		fmt.Printf("Choose file (%s)?\n> ", path)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(input)
		return string(getContent(input))

	} else {
		return string(getContent(path))

	}
}

func getContent(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
