package main

import "fmt"

func main() {

	addition(10, 10)
	subtraction(20, 10)
	multiplication(2, 10)
	division(10, 2)

	fmt.Println("Grouping Codes in functions")

	fmt.Println(addition(10, 10))

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

func multiplication(num_1, num_2 int) int {

	sum := num_1 * num_2

	fmt.Printf("MULTIPLICATION  = %d", sum)

	return sum

}
func division(num_1, num_2 int) int {

	sum := num_1 / num_2

	fmt.Printf("Division  = %d", sum)

	return sum

}
