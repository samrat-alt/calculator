package main

import (
	"calculator/sum"
	"fmt"
)

func main() {
	var num1, num2 float64

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	fmt.Println("Choose operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multilpy")
	fmt.Println("4. Division")

	var choice int
	fmt.Scanln(&choice)

	var result float64
	var operation string
	switch choice {
	case 1:
		result = sum.Add(num1, num2)
		operation = "Addition"
	case 2:
		result = sum.Sub(num1, num2)
		operation = "Subtraction"

	case 3:
		result = sum.Multiply(num1, num2)
		operation = "Multilpy"

	case 4:
		result = sum.Division(num1, num2)
		operation = "Division"
	default:
		fmt.Println("Invalid choice")
		return
	}
	fmt.Println(result, operation)
}
