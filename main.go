package main

import (
	"fmt"
)

type byteMachine struct {
	memory   []byte
	stack    []int
	pointer  int
	register [3]int
}

var Bigbytemachine byteMachine

func pop(someStack []int) []int {
	return someStack[:len(someStack)-1]
}
func push(someStack []int, otherStack []int) []int {
	return append(someStack, otherStack...)
}

// func store(someregister [3]int, value int, index int) [3]int {

// }

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
			stack = pop(stack)
			stack = pop(stack)
			stack = push(stack, []int{addedStack})
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf8 { //Subtraction feature
			stack := Bigbytemachine.stack
			subtractedStack := stack[len(stack)-1] - stack[len(stack)-2]
			stack = pop(stack)
			stack = pop(stack)
			stack = push(stack, []int{subtractedStack})
			//stack = stack[:len(stack)-2]
			//stack = append(stack, subtractedStack)
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf7 { //Division feature
			stack := Bigbytemachine.stack
			dividedStack := stack[len(stack)-2] / stack[len(stack)-1]
			stack = pop(stack)
			stack = pop(stack)
			stack = push(stack, []int{dividedStack})
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf6 { //Modulo feature
			stack := Bigbytemachine.stack
			ModuloStack := stack[len(stack)-2] % stack[len(stack)-1]
			stack = pop(stack)
			stack = pop(stack)
			stack = push(stack, []int{ModuloStack})
			Bigbytemachine.stack = stack
			Bigbytemachine.pointer++
			continue
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf5 { //store feature
			index := Bigbytemachine.memory[Bigbytemachine.pointer+1]
			value := Bigbytemachine.stack[len(Bigbytemachine.stack)-1]
			Bigbytemachine.register[index] = value
			Bigbytemachine.stack = pop(Bigbytemachine.stack)
			Bigbytemachine.pointer += 2
			fmt.Println("Register: ", Bigbytemachine.register)
		}
		if Bigbytemachine.memory[Bigbytemachine.pointer] == 0xf4 { //load feature
			index := Bigbytemachine.memory[Bigbytemachine.pointer+1]
			value := Bigbytemachine.register[index]
			Bigbytemachine.stack = append(Bigbytemachine.stack, value)
			Bigbytemachine.pointer += 2
			continue
		}
	}
	fmt.Println("Stack: ", Bigbytemachine.stack)

}
func main() {
	steps := []byte{0xff, 0x01, 0xf5, 0x00, 0xf4, 0x00}
	memorySteps(steps)
}
