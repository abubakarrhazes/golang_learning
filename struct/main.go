package main

import "fmt"

func main() {
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
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

}

type point struct {
	x float32
	y float32
	z float32
}
