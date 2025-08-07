package main

import (
	"fmt"
)

type byteMachine struct {
	memory  []byte
	stack   []int
	pointer int
}

var Bigbytemachine byteMachine

func memorySteps(steps []byte) {
	Bigbytemachine.memory = steps
	Bigbytemachine.pointer = 0
	for Bigbytemachine.pointer < len(Bigbytemachine.memory) { // Push Feature
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xff {
			nextOne := Bigbytemachine.memory[Bigbytemachine.pointer+1]
			Bigbytemachine.stack = append(Bigbytemachine.stack, int(nextOne))
			Bigbytemachine.pointer += 2
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf9 { // Addition feature
			stack := Bigbytemachine.stack
			addedStack := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, addedStack)
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf8 { //Subtraction feature
			stack := Bigbytemachine.stack
			subtractedStack := stack[len(stack)-1] - stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, subtractedStack)
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf7 { //Division feature
			stack := Bigbytemachine.stack
			dividedStack := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, dividedStack)
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf6 { //Modulo feature
			stack := Bigbytemachine.stack
			ModuloStack := stack[len(stack)-2] % stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, ModuloStack)
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
	}
	fmt.Println("Stack: ", Bigbytemachine.stack)
}
func main() {
	steps := []byte{0xff, 0x01, 0xff, 0x04, 0xf6}
	memorySteps(steps)
}
