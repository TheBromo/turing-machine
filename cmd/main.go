package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	logger "github.com/thebromo/turing-machine/log"
	parser "github.com/thebromo/turing-machine/parser"
)

func main() {
	fmt.Print("Stepmpdus (y/N)?\n> ")
	stepModus := askBool()
	file := askFile()

	//read file
	input := file
	re := regexp.MustCompile("111")
	split := re.Split(input, 2)

	if len(split) == 2 {
		tape := parser.InitTape(split[1])
		turing, err := parser.InitMachine(split[0], tape)

		if err != nil {
			panic(err)
		}
		count := 0
		for err == nil {
			count++

			err = turing.DoStep()
			if stepModus {
				fmt.Println("> " + strconv.Itoa(count))
				logger.PrintMachine(turing, err)
				time.Sleep(2 * time.Second) // pauses execution for 2 seconds
			}
		}
		fmt.Println("\n> " + strconv.Itoa(count) + "\n")
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
