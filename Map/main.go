package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Shree Ganeshai Namah!")

	var firstName string
	var lastName string
	var email string
	var noOfFruits uint
	

	//want to make a list of maps

	var users = make([]map[string]string, 0)  //10 is initial size

	for i:= 1; i<3; i++{
		fmt.Printf("Details for user %v: \n", i)

		fmt.Printf("Enter your First Name: \n");
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: \n");
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: \n");
		fmt.Scan(&email)
		fmt.Printf("Enter the number of fruits you want: \n");
		fmt.Scan(&noOfFruits)

		var userData = make(map[string]string, 4)  //data type is map[string]string
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["noOfFruits"] = strconv.FormatUint(uint64(noOfFruits), 10)  //typecasting to uint64, because noOffruits declared above in an uint and the first argument of the FormatUint is uint64
		
		users = append(users, userData)
	}

	fmt.Printf("%T \n", users)
	fmt.Println(users[0]);
	fmt.Println(users[1]);

}

/*
suppose you want to store like the following:
firstName: Praveen
lastName: Kumar
Email: praveen08h@gmail.com

*/

/*
	It maps unique keys to values.

	All keys have the same data Type and all the values have the same data Type.

	A map element is not a variable, and we cannot take its address.

	strconv converts strings to other datatypes :is short for string conversion
		: https://pkg.go.dev/strconv#pkg-overview  : refer for more
		: strconv.Uint(integer-value , 10) : 10 here is base, coverts the decimal unsigned integer value to string

*/