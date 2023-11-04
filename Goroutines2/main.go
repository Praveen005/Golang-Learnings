package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
func main(){

	wg.Add(1)
	go greeter(2)


	//you can only use range with collections like arrays, slices, maps, or channels. val1 is int so will give an error
	/*
		for _, v := range val1{
			fmt.Println(v)
		}
	*/

	for i := 1; i<5; i++{
		fmt.Println(i)
	}

	wg.Wait()
}

func greeter(d time.Duration){
	time.Sleep(d * time.Second)
	fmt.Printf("Hello Praveen Ji!\n")

	wg.Done()
}
/*
If you just run the code, it will only print
1
2
3
4
and not he greeting message. why?
when the execution of main thread is complete, execution of the program terminates, which is wrong
It is not waiting for the another goroutine to complete.

so, we need to tell the main thread to wait till the another thread is done doing its job. for that we will make wait group

wait group:
	waits for the launched goroutine to finish
	package sync provides basic synchronization functionality
	WaitGroup has 3 functions that we can call using 'wg' variable above
		1. Add: sets the number of goroutines to wait for(increases the counter by provided number)
		2. wait: Blocks untill the WaitGroup counter is 0
		3. Decrements the WaitGroup counter by 1, this is called by the goroutine to indicate that it is finished
*/

/*
You get concurrency in Java as well, so why Go?
In Java:
writting concurrent code is more complex
more overhead

creating threads is more expensive, slow startup time
Heavyweight and need more hardware resources

It uses operating system thread

managed by kernel

Are hardware dependent

No easy way to communicate between the threads

In Go:
Go is using, what's called green thread

Green Thread is actually an abstraction of an actual thread, which is an opearting system thread, called a goroutine

Managed by the go runtime, we are only interacting with these high level goroutines

Cheaper & Lightweight

you can run hundred and thouands of goroutines without affecting the performance

Goroutines have channels to communicate between the threads : it is a built-in functionality

*/