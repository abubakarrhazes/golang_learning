package main

import "fmt"

func main() {
	var pt1 point

	var user User

	var address Address
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2

	// Declaring User Variables Here

	user.FirstName = "Abubakar"
	user.LastName = "Nuuman"
	user.OtherName = "Adam"

	// Declaring Address Varibales Here

	address.Country = "Nigeria"
	address.State = "Sokoto"
	address.LGA = "Kware"
	address.Ward = "Durbawa"
	address.HomeAddress = "Durbawa Town "
	/*
		Consider You want to store  location of point as X and Y Coordinate



	*/
	// pt1X := 3.1
	// pt1Y := 5.7

	// pt2X := 5.6
	// pt2Y := 3.8
	fmt.Println("Learning GO Struct")
	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)
	fmt.Println("")

}

type point struct {
	x float32
	y float32
	z float32
}

type User struct {
	FirstName   string
	LastName    string
	OtherName   string
	DateOfBirth string
}

type Address struct {
	Country     string
	State       string
	LGA         string
	Ward        string
	HomeAddress string
}
