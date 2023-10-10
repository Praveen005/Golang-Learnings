package main

import(
	"fmt"
)

func main(){
	// var bookings= [50]string{"Prateek", "Praveen", "Pratyush"}  //only the same type of data can be stored, size is fixed
	var bookings [50]string //the type of this string is [50]string
	bookings[0]= "Prateek"
	bookings[1]= "Praveen"

	fmt.Println(bookings[0])
	fmt.Println(bookings[1])
	fmt.Println(bookings)
	fmt.Printf("The complete array: %v %T\n",bookings, bookings) 
	fmt.Printf("The length of the array is %v\n", len(bookings))  //len(array) gives the length of the array

	var bookings2 []string  //this is a slice
	bookings2 = append(bookings2, "Mango")
	fmt.Println(bookings2)

	slice := make([]int, 3, 5) // length: 3, capacity: 5
	fmt.Println(len(slice)); //length
	fmt.Println(cap(slice)); //capacity

	slice2 := []int{1,2,3}; //slice initialized
	fmt.Println(slice2)
	slice2= append(slice2, 4,5)
	fmt.Println(slice2)
	slice2[3]= 34;
	fmt.Println(slice2)

	var slice3 = []int{} //empty slice
	fmt.Println(len(slice3))

	slice4 := []int{}  //empty slice
	fmt.Println(len(slice4))

	//Adding one slice to another
	slice5 := []int{}
	slice5 = append(slice5, slice4...)
	fmt.Println(slice5)
	//here ... is known as the "variadic" or "unpacking" operator.It is used to unpack the elements of a slice and pass them as separate arguments to a variadic function.

	// The ... is used after slice4, and it tells Go to unpack the elements of slice4 and pass them as individual arguments to the append function. Without ..., you would get a compilation error because you cannot directly append one slice to another without unpacking its elements.

	// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function. Just a generic name

	//copying one slice to another
	slice6 := append([]int(nil), slice5...)
	fmt.Println(slice6)

	// In Go, nil is a special value that represents a zero value
	// []int(nil) creates an empty slice of type int with a length and capacity of zero. This is essentially creating a new empty slice.
	
}

// Remember, the slice header is always updated by a call to append, cz it starts pointing to a new array with new capacity :)

/*
There is one issue with Array, what if you don't know the size of the array beforehand?
: This is where slice comes in.

slices are flexible and Powerful: variable length or get an sub array of your own
slices are index based and have size, but is resized when needed
To add elements we use append(), it adds element to the end, grows the size if needed and returns the updated slice

*/
/*

If you pass slices to a function by value, you know deep down slice is a struct, which contains the pointer to first element of an array and a length variable
and inside that function you change the value, it will be reflected back in the original slice. why because it is ltimately pointing to an array and values at that address is being chnaged.

Now if you reduce the length as well inside that function, it will not be reflected inside the original slice, it may point out to the orignal array but ultimately it is a new slice.const

so, if you want to alter the length of the the original slice, pass it by reference

read for more: https://go.dev/blog/slices

*/