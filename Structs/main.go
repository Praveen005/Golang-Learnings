package main

import(
	"fmt"
	"time"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Shree Ganeshay Namah 🙏")

	type Employee struct {
			ID int
			Name string
			Address string
			DoB time.Time
			Position string
			Salary int
			ManagerID int
		}
		var googleEmployee Employee

	// var employeeData = make([]Employee, 10) //declaring a splice of structs

	var employeeData []Employee
	
	for i := 1; i<2; i++{
		fmt.Printf("Enter details for User %d: \n", i)
		fmt.Println("ID:")
		fmt.Scanln(&googleEmployee.ID)

		fmt.Println("Name:")
		fmt.Scanln(&googleEmployee.Name)
		// Clear the input buffer
		// var dummy string
		// fmt.Scanln(&dummy)

		fmt.Println("Address:")
		//reads the entire line, including spaces
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			googleEmployee.Address = scanner.Text()
		}

		fmt.Println("Position:")
		fmt.Scan(&googleEmployee.Position)

		fmt.Println("Salary:")
		fmt.Scan(&googleEmployee.Salary)

		fmt.Println("ManagerID:")
		fmt.Scan(&googleEmployee.ManagerID)

		googleEmployee.DoB= time.Now()

		employeeData = append(employeeData, googleEmployee)
	}
	
	/*
	var t time.Time
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Now().In(loc)

	fmt.Println(t)
	*/

	fmt.Println(employeeData[0])
}

/*
	structs:  stands for structure
	
	collects different types of data

	formal def: (The Go Programming Language Book: ) :  struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity. Each value is called a field.

	you remember in maps, there was a restriction to use a single data type, var myMap= make(map[string]int, 10)

	All of these fields are collected into a single entity that can be copied as a unit, passed to functions and returned by them, stored in arrays, and so on.

*/