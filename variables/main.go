package main

import "fmt"

func main() {

	//We can declared a variable using var keyword and also by assigning the data type
	var num int = 10 // 0

	var name string = "Abubakar Nuuman Adam" // ""

	var loading bool = true // Default value when not assign will false

	fmt.Printf(" VARIABLES DELCLARED WITH VAR KEYWORD %d", num)

	fmt.Printf(" VARIABLES DELCLARED WITH VAR KEYWORD %s", name)

	fmt.Printf(" VARIABLES DELCLARED WITH VAR KEYWORD %t", loading)

	// Declaring Variables using  Inline and short-hand declaration statement

	var num1, num2 int = 10, 20

	animal := "DOG"

	underAge := false

	fmt.Println(underAge)

	fmt.Println("Animal here is " + animal)

	fmt.Printf(" VARIABLES DELCLARED WITH Inline Method %d  and %d", num1, num2)

	fmt.Println("Learning Variables in Golang")

}
