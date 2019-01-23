// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 12B:
//    Použití obousměrně vázaného seznamu ve funkci zásobníku.

package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Stack list.List

func printStack(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%2d ", e.Value)
	}
	println()
}

func push(l *list.List, number int) {
	l.PushBack(number)
}

func pop(l *list.List) int {
	tos := l.Back()
	l.Remove(tos)
	return tos.Value.(int)
}

func main() {
	expression := "1 2 + 2 3 * 8 + *"
	terms := strings.Split(expression, " ")
	stack := list.New()

	for _, term := range terms {
		switch term {
		case "+":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand1+operand2)
			print("+ :     ")
			printStack(stack)
		case "-":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2-operand1)
			print("- :     ")
			printStack(stack)
		case "*":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand1*operand2)
			print("* :     ")
			printStack(stack)
		case "/":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2/operand1)
			print("/ :     ")
			printStack(stack)
		default:
			number, err := strconv.Atoi(term)
			if err == nil {
				push(stack, number)
			}
			fmt.Printf("%-2d:     ", number)
			printStack(stack)
		}
	}
	print("Result: ")
	printStack(stack)
}
