package main

//When you import a package but not use it, error is thrown
import (
	"fmt"
)

//Print() : doesn't print a new-line at the end
//Println() : prints a new line at end
//when you declare a variable and assign it a value immediately, it won't throw an error, it will automatically understand from the value assigned, but when declared but not assigned, you need to provide the data type

func main() {

	conferenceName := "Go Conference" // := go will understand the data type, after it gets assigned a value
	const conferenceTickets int = 50  // := can't use with const
	var remainingTickets uint = 50    //unsigned int

	//%T prints the datatype
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)                                    //space around confereName variable automatically appears in result
	fmt.Println("We have", conferenceTickets, "tickets in total and", remainingTickets, "tickets available") //%v : variable reference
	fmt.Println("Get your tickets here to attend")
	//%v is default

	//when you declare a variable and assign it a value immediately, it won't throw an error, it will automatically understand from the value assigned, but when declared but not assigned, you need to provide the data type

	var userName = "Tom"
	/*
	why getting warning below?
	ans:
	Not recomended:
	var userTickets int
	userTickets = 2

	recomended:
	var userTickets int = 2
	or
	userTickets := 2;
	*/
	var userTickets int
	userTickets = 2
	fmt.Printf("User %v got %v tickets\n", userName, userTickets)

	//user- Input
	var userName2 string
	fmt.Scan(&userName2) //&userName2 is pointer, passing the variable's address
	fmt.Println(userName2)

	var data= 34
	// %T is the place holder for type of variable
	fmt.Printf("The data type of %v is %T", data, data)


	str := "Hi, Bunny"
	runes1 := []rune(str) //stores the unicode values of the characters
	fmt.Println(runes1)  //prints the decimal values of the unicodes

	runes2 := []rune{72, 105, 44, 32, 66, 117, 110, 110, 121}
	fmt.Println(string(runes2)) //type converts the rune to a string and prints, "Hi, Bunny"

}



/*
Before running the go file, you need to initialize the folder
	go mod init Boking-App
	-> It will generate a go.mod file


In Go, everything is organised in a package.

The main fucntion is the entry point of the Go Program

A go program can have only one main fucntion. Cz you only can have 1 entry point.

Go's standard library, provides different core packages for us to use,
'fmt' is one such package, which you can use by importing it.

Think of packages a container which contains several fucntionalities

Variable name once declared must be used.
Similarly a package once declared must also be used.


Data Types: Strings Integers Booleans Maps Arrays

Go is a statically typed language: you need to tell the compiler, the data type when declaring the variable

one question you might have is, when we do,
	var value= 78
we are not declaring the data type, why no errors here?

ans: when we declare and assign immediately, Go can infer the type of value you assign itself. so when you don't initialize it, you need to specify the data type


You get errors durng compile time and not runtime	.const
Go to have, uint8, uint16, uint32, uint64, int8, int16, int32, int64
floating types: float32, float64, complex64, complex128


remember, using := you can't declare constants


%s for strings, %d for integers, %f for floating-point numbers
*/

/*
A pointer is avariable that points to the memory address of another variable.

var tickets = 50, say memory address is 0xc23456

pointer &tickets = 0xc300002 which stores 0xc23456 and 0xc23456 points to 50

When you calculation in go, they need to have same type, ex: int-uint will give error,
you either, make them both of the same datatype or convert one to the another's datatype
*/
