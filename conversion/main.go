package main

import (
	"fmt"
)


func main(){
	const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
	const b = 15 / 4           // b == 3     (untyped integer constant)
	const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
	const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
	const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
	const d = 1 << 3.0         // d == 8     (untyped integer constant)
	const e = 1.0 << 3         // e == 8     (untyped integer constant)
	//const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
	//const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
	const h = "foo" > "bar"    // h == true  (untyped boolean constant)
	const j = true             // j == true  (untyped boolean constant)
	const k = 'w' + 1          // k == 'x'   (untyped rune constant)
	const l = "hi"             // l == "hi"  (untyped string constant)
	const m = string(k)        // m == "x"   (type string)
	const Σ = 1 - 0.707i       //            (untyped complex constant)
	const Δ = Σ + 2.0e-4       //            (untyped complex constant)

	const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
	const Four int8 = Huge >> 98  // Four == 4                                (type int8)

	// fmt.Println(Huge)   //will give error, as it oerflows



	// Note: This form of conversion may eventually be removed from the language. The go vet tool flags certain integer-to-string conversions as potential errors. Library functions such as utf8.AppendRune or utf8.EncodeRune should be used instead.
	fmt.Println(string('a') )         // "a"
	fmt.Println(string(65))          // "A"
	fmt.Println(string('\xf8') )      // "\u00f8" == "ø" == "\xc3\xb8"
	fmt.Println(string(-1))          // "\ufffd" == "\xef\xbf\xbd"




	// Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.
	//will print the decimal representation of the unicodes
	fmt.Println([]rune(string("白鵬翔")))   // []rune{0x767d, 0x9d6c, 0x7fd4}
	fmt.Println([]rune(""))                  // []rune{}

	fmt.Println([]rune("白鵬翔"))              // []rune{0x767d, 0x9d6c, 0x7fd4}

	fmt.Println([]rune("♫♬"))              // []myRune{0x266b, 0x266c}
	fmt.Println([]rune(string("🌐")))    // []myRune{0x1f310}



	// When converting a floating-point number to an integer, the fraction is discarded (truncation towards zero).
	
	// fmt.Println(int(1.2)) // cannot convert 1.2 (untyped float constant) to type int
}
//https://go.dev/ref/spec#Conversions

/*
package main

import (
    "fmt"
    "math/big"
)

func main() {
    Huge := new(big.Int)
    Huge.SetString("1267650600228229401496703205376", 10)
    fmt.Println(Huge)
}


*/