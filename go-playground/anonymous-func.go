package main

import (
	"fmt"
	"strconv"
)

type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": multiply,
	"/": divide,
}

func main() {
	fmt.Println("Anonymous")

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "+", "3"},
		{"two", "+", "three"},
		{"two", "zero", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression: ", expression)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operation: ", op)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Printf("Result: %d\n", result)
	}
}

func add(i int, j int) int      { return i + j }
func sub(i int, j int) int      { return i - j }
func multiply(i int, j int) int { return i * j }
func divide(i int, j int) int   { return i / j }
