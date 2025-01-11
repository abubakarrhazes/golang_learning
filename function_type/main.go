package main

import "fmt"

func main() {

	addition(10, 10)
	subtraction(20, 10)

	fmt.Println("Grouping Codes in functions")

}

// METHODS THAT ADD TWO INTEGER NUMBERS
func addition(num_1, num_2 int) int {
	sum := num_1 + num_2

	fmt.Printf("ADDITION = %d", sum)

	return sum

}

func subtraction(num_1, num_2 int) int {

	sum := num_1 - num_2

	fmt.Printf("SUBTRACTION  = %d", sum)

	return sum

}
