package main

import (
	"fmt"
)

type byteMachine struct {
	memory []byte
	stack  []int
	//pointer int
}

var Bigbytemachine byteMachine

func memorySteps(steps []byte) {
	Bigbytemachine.memory = steps
	for i := 0; i < len(Bigbytemachine.memory); i++ {
		if Bigbytemachine.memory[i] == 0x00 {
			nextOne := Bigbytemachine.memory[i+1]
			Bigbytemachine.stack = append(Bigbytemachine.stack, int(nextOne))
			i++
		}
		if Bigbytemachine.memory[i] == 0x01 {
			stack := Bigbytemachine.stack
			addedStack := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, addedStack)
			Bigbytemachine.stack = stack
		}
	}
	fmt.Println("Stack: ", Bigbytemachine.stack)
}
func main() {
	steps := []byte{0x00, 0x04, 0x00, 0x05, 0x01}
	memorySteps(steps)
}
