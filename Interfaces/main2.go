package main

import (
	"math"
	"fmt"
)

//when you call this interface, it should give you both the Perimeter as well as the Area
type shape interface{
	area() float64
	perimeter() float64
}

//on what type of shapes do you want the area/perimeter to be computed
type rect struct{
	len int
	width int
}

type circle struct{
	radius int
}

//Defining methods, mind you method is a special type of function
/* In between the keyword func and the name of the function, we’ve added a receiver.
The receiver is like a parameter—it has a name and a type—but by creating the func‐
tion in this way, it allows us to call the function using the . operator:
fmt.Println(c.area())
*/
func (r rect) area() float64{
	return float64(r.len) * float64(r.width)
} 

func (c circle) area() float64{
	return math.Pi* float64(c.radius)* float64(c.radius)
}

func (r rect) perimeter() float64{
	return 2*(float64(r.len) + float64(r.width))
} 

func (c circle) perimeter() float64{
	return 2* math.Pi* float64(c.radius)
}


func measure(s shape){
	fmt.Println(s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}