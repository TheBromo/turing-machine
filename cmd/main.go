package main

import "github.com/thebromo/turing-machine/machine"

func main() {
	print("Hello world")

	machine := machine.Machine{
		Name: "test",
	}

	print(machine.GetTestString())
}
