package main

import "fmt"

func main() {

	var num1 int
	print("Enter first number: ")
	fmt.Scan(&num1)

	var num2 int
	print("Enter second number: ")
	fmt.Scan(&num2)

	add := num1 + num2

	sub := num1 - num2

	mul := num1 * num2

	div := num1 / num2

	print("The addition is: ", add, "\n")

	print("The subtraction is: ", sub, "\n")

	print("The multiplication is: ", mul, "\n")
	
	print("The division is: ", div, "\n")

}
