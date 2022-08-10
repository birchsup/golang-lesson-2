package main

import (
	"fmt"
	"math"
)

func main() {

	circumference()

}

func amount() {

	var S int
	fmt.Println("type a number")
	fmt.Scanf("%d\n", &S)

	fmt.Println("number of hundreds = ", S/100)
	fmt.Println("number of tens = ", S/10%10)
	fmt.Println("number of units = ", S%10)

}

func area_of_a_rectangle() {

	var a int
	fmt.Println("type a")
	fmt.Scanf("%d\n", &a)

	var b int
	fmt.Println("type b")
	fmt.Scanf("%d\n", &b)

	fmt.Println("area of a rectangle = ", a*b)

}

func circumference() {

	var S float64
	fmt.Println("type area of a circle")
	fmt.Scanf("%f\n", &S)

	res := math.Sqrt(S)
	pi := math.Pi
	t := 2 * res
	diametr := t / pi
	circumference := pi * diametr

	fmt.Println("circle diameter = ", diametr)
	fmt.Println("circumference = ", circumference)

}
