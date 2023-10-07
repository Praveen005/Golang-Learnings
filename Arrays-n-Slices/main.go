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
	
}

/*
There is one issue with Array, what if you don't know the size of the array beforehand?
: This is where slice comes in.

slices are flexible and Powerful: variable length or get an sub array of your own
slices are index based and have size, but is resized when needed
To add elements we use append(), it adds element to the end, grows the size if needed and returns the updated slice

*/