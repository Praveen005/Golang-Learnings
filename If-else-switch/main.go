package main

import (
	"fmt"
	"strings"
)

func main(){
	var i int= 1

	for{
		if i<5{
			fmt.Println(i)
		}else if i>= 5 && i<10{
			fmt.Println(i);
		}else{   //don't put else or else if in a seperate line
			break
		}
		i++
	}


	//User Input Validation

	var firstName string
	fmt.Scan(&firstName)
	var lastName string
	fmt.Scan(&lastName)
	var emailAdd string
	fmt.Scan(&emailAdd)

	//rules: first name and last name should atleast be two characters
	//can use len() built-in function

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	if(isValidName){
		fmt.Println("User has entered a valid name")
	}else{
		fmt.Println("Not a valid user name!")
	}

	//email validation
	//check if the string contains a '@' symbol
	//will use strings.Contains(str, "@")

	isValidEmail := strings.Contains(emailAdd, "@")
	if(isValidEmail){
		fmt.Println("User has entered a valid email address.")
	}else{
		fmt.Println("Invalid email entered!")
	}

	if(isValidEmail && isValidName){
		fmt.Println("Valid User Entry!!")
	}





	//Switch statements

	var cityName string
	fmt.Scan(&cityName)

	switch cityName{
		case "Singapore" :
			fmt.Println("You will be served Sushi.")
			
		case "London" :
			fmt.Println("You will be served Hot Dog.")
			
		case "New Delhi" :
			fmt.Println("You will be served Chola Bhatura.")

		default :
			fmt.Println("You will be served Maggi!")
	}
}