package main

import(
	"fmt"
)
//Top level Variables
// :defines at the top outside all the functions
// : They can be accessed inside any of te functions
// : And in all the files that are in the same package
// : Package level variables can't be created using := 
// : Best practice is to define variables locally as much as possible

func main(){
	fmt.Println("Shree Ganeshay Namah!!")

	val1, val2, val3 := greet("Prateek")

	fmt.Println(val1, val2, val3)

	
}

func greet(name string)  (int, int, int) {
	fmt.Println("Welcome to our space,", name)
	return 1, 2, 3
}

//In Golang you can return any number of values from a function, also define the datatypes of all the return values by enclosing them in parenthesis

// In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.