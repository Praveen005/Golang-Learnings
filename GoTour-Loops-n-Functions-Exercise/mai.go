package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var res float64
	var z float64 = 1.0
	var i int = 1
	for i<1000{
		if math.Abs(x - z*z) < 0.000001 {
			return res;
		}
		res = z
        z = (z + x/z)/2
		i++
	}
	return res
}

func main() {
	fmt.Println(Sqrt(21))
}
