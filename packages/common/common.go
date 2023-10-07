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



//we made out own package named "common" , for that made a folder named "common"
//to use the methos/functions defined here, we need to export it from here
//to export a function, we just capitalize the functiuon name. why? remember we were usig Print from fmt, where P is capital
//you can also export variables by capitalizing the variable names

