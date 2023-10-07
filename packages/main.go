package main

import (
	"fmt"
	"packages/common"   //this is how we import user defined package
)

var value int =43
var Value2 int =434  //exported Value2 by capitalizing the first letter

func main(){
	fmt.Println("Shri Ganeshay Namah!!");
	// can we modularize the application? 
	//can the application have multiple files, with each module tackling one task
	//: Go is organised into packages
	//: A package is a collection of one file or several files
	//: Like we just saw, that we had a main.go file with package main, we can have multiple files within package main

	helperFunc()
	common.CommonFunc(Value2) //this is how we access the methods defined in package "common"
	fmt.Println("Printing the value the package level variable declared in pacage 'common' :", common.MyValue)
	fmt.Println(packageVariable)
}

var packageVariable int = 53

//scopes

//3 levels iof scope : Loca Package & Global

//you saw how we define local variables

//package variables : var value = 32   //see it is defined out side the functions and see, all are small(koi zaruri nahi hai)

//agar ye jo package variables ain, if the first letter is capital then it is global variable, like: var Value = 32