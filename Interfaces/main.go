package main

import "fmt"

// Define an interface named Writer with a single method Write.
type Writer interface {
    Write([]byte) (int, error)
}

// Implement the Writer interface for a custom type MyWriter.
type MyWriter struct{}

func (w MyWriter) Write(data []byte) (int, error) {
    // Implement the Write method for MyWriter.
    // Typically, this method would write data to some output.
    fmt.Println(string(data))
    return len(data), nil
}

type circle2 struct{
	x,y,z int
}

func modifyCircle(c *circle2){
	c.x ++
	c.y ++
	c.z++
}

func main() {
    // Create an instance of MyWriter.
    myWriter := MyWriter{}

    // Use the interface type as a parameter.
    writeSomething(myWriter)

	r := rect{len:3 , width: 5}
	c := circle{radius: 7}

	measure(r)
	measure(c)

	var c2 circle2;
	c2.x= 10
	c2.y= 12
	c2.z= 23

	fmt.Println(c2)
	modifyCircle(&c2)
	fmt.Println(c2)


}

func writeSomething(w Writer) {
    data := []byte("Hello, Golang!")
    w.Write(data)
}




/*
	Interfaces are collections of method signatures.

	A type "implements" an interface if it has all of the methods of the given interface defined on it.

	****************************************************************************************************
		func (w MyWriter) Write(data []byte) (int, error) {
			fmt.Println(string(data))
			return len(data), nil
		}

		what is the name of this function?
		Answer:
			The name of the function you provided is "Write," and it's a method of the `MyWriter` type. In Go, a function that has a receiver (in this case, `(w MyWriter)`) is a method. The function name, `Write` in this case, is the method name, and it is associated with the `MyWriter` type. This method is used to satisfy the `Writer` interface, as it has the same method signature (name, parameters, and return types) as the `Write` method defined in the `Writer` interface.
	**********************************************************************************************************
*/