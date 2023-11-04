package main

import "fmt"


func helperFunc(){
	fmt.Println(value) // we can access and use the variables defined at package level in here, but have to use this function in main.go
}

//Also you would have notices that you can't define the func main(){} here again.