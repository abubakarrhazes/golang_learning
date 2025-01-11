package main

import "fmt"

func main() {

	// In Golang We have 3 control/condition statement

	/*

		1: IF/ELSE STATEMENT
		2: SWITCH STATEMENT
		3: SELECT STATEMENT

	*/

	//This program check if a user is eligble to drive
	num := 1
	checkEligibility := num >= 18
	if checkEligibility {
		fmt.Println("YOU ARE ELIGLBLE TO DRIVE MAN :D")

	} else {

		fmt.Println("NAH YOU ARE STILL YOUNG")

	}

	divisionOfNumber := 10

	checkDivision := divisionOfNumber%2 == 0

	if checkDivision {
		fmt.Println("Divisible ")

	} else {
		fmt.Println("Not Divisible Man")
	}

	fmt.Println("Conditional Statements In Golang")

	// TESTING THE FUNCTONS IN THE MAIN GO FILE USING STATEMENT CONDITIONS

	weather := false

	if isRaining(weather) || isSnowing(true) {

		fmt.Println("NOT A PERFECT WEATHER MAN STAY INDOOR  ")

	} else {
		fmt.Println("NOT RAINY ")
	}

	v, err := doSomething()

	if err {
		fmt.Println("Error Occured")
	} else {
		fmt.Println(v)
	}

}

func doSomething() (int, bool) {

	return 5, false
}

// CHECKING CLOUDY FUNCTION

func isCloudy() bool {

	fmt.Println("IT CLOUDY MAN")

	return true

}

//CHECKING RAINING FUNCTION

func isRaining(condition bool) bool {

	//fmt.Println("IT RAINING MAN , SO STAY INDOORS")

	return condition

}

//CHECKING THE SNOWING FUNCTION

func isSnowing(condition bool) bool {

	// fmt.Println("IT SNOWING MAN , SO STAY INDOORS ")

	return condition

}
