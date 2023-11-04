package main

import(
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)  //run concurrently to the fib()
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) { //used for making a spinning animation till the function fib() calculates the value concurrently
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)  //\r is a carriage return character that moves the cursor to the beginning of the current line. This allows the animation to overwrite itself, creating the appearance of spinning.
			time.Sleep(delay)  // prints r(one character, say '-' and waits for the 'delay' period)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
/*
	After several seconds of animation, the fib(45) call returns and the main function prints its
	result:
	Fibonacci(45) = 1134903170

	The main function then returns. When this happens, all goroutines are abruptly terminated
	and the program exits. Other than by returning from main or exiting the program, there is no
	programmatic way for one goroutine to stop another, but there are ways to
	communicate with a goroutine to request that it stop itself.

*/
	
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

/*
	A defer statement defers the execution of a function until the surrounding function returns.

	The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

	refer for more: https://go.dev/blog/defer-panic-and-recover

*/