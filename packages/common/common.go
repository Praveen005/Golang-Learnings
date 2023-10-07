package common

import (
	"fmt"
)

var MyValue int

func CommonFunc(Value2 int) {
	fmt.Println(Value2)
}

func init(){
	MyValue = 768
}



//we made our own package named "common" , for that made a folder named "common"
//to use the methods/functions defined here, we need to export it from here
//to export a function, we just capitalize the function name. why? remember, when we use Print from fmt,  'P' is capital in it
//you can also export variables by capitalizing the variable names

