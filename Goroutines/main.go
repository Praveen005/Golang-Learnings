package main

import(
	"fmt"
)

func main(){
	fmt.Println("Shree Ganeshay Namah 🙏")
	var p int = 32
	var value = fmt.Sprintf("This is a greeting message %d\n", p)
	fmt.Println(value)
}
/*
	fmt.Sprintf is a function in Go's fmt package that is used for string formatting. It allows you to create a formatted string by combining values with format specifiers
*/

/*
	In Go, each concurrently executing activity is called a goroutine.

	When a program starts, its only goroutine is the one that calls the main func tion, so we call it
	the main goroutine. New goroutines are created by the go statement. Syntactically, a go statement is an ordinary function or method call prefixed by the key word go. A go statement causes the function to be calle d in a newly created goroutine.

	f() // call f(); wait for it to return
	go f() // create a new goroutine that calls f(); don't wait


*/