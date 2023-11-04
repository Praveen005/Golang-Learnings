package main

import (
	"fmt"
)
/*
const Spanish ="Spanish"
const French ="French"
const Englishprefix ="Hello, "
const Spanishprefix ="Hola, "
const Frenchprefix ="Bonjour, "
*/

const(
	Spanish ="Spanish"
	French ="French"
	Englishprefix ="Hello, "
	Spanishprefix ="Hola, "
	Frenchprefix ="Bonjour, "
)

/*
func hello(name string, language string) string{
	if(name == ""){
		name= "world"
	}

	if(language == Spanish){
		return Spanishprefix + name
	}else if(language == French){
		return Frenchprefix + name
	}
	return Englishprefix + name
}
*/

func hello(name, language string) string{
	if name == ""{
		name= "world"
	}

	prefix := Englishprefix

	switch language{
		case Spanish:
			prefix = Spanishprefix
		case French:
			prefix = Frenchprefix
	}

	return prefix + name
}

func main(){
	fmt.Println(hello("Praveen", "Spanish"))
}



/*
	To run the program: 				go run .   or  go run main.go
	To create the mod file: 			go mod init folder-name(Hello-World here)
	Go add the folder to the go.work:   ./TDD/Hello-World


	In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.

	 we can group constants in a block instead of declaring them each on their own line.
*/