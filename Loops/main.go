package main

import (
	"fmt"
	"strings"
)


func main(){
	fmt.Println("Shri Ganeshay Namah!!");


	for i := 0; i < 5; i++ {  //one way
		// Loop body
		fmt.Println(i)
	}

	var i int =0;
	for {
		// Infinite loop
		// You would use a break statement to exit when needed
		if(i>10){
			break;
		}
		fmt.Printf("%v ", i)
		i++;
	}
	fmt.Printf("\n")
	
	

	numbers := []int{1, 2, 3, 4, 5}
	for index1, value1 := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index1, value1)
	}

	//want to split a string about the white space?
	str := "Today is Saturday"
	var tokens []string  //declaring a splice
	tokens= strings.Fields(str)   //storing the splitted tokens in the splice
	for index, val := range tokens{
		fmt.Println(index, val)
	}

	//if you don't want to use index, then use _ (underscore) to ignore it, else you will get error, "declared but not used"
	for _, value := range tokens{
		fmt.Println(value)
	}


}

/*
Range iterates over elememnts for different data structures(not only arrays or slices)
: It provides index and value for each element

strings.Fields() splits with white space as seperator


*/