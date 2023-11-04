package main 

import (
	"fmt"
)

func main(){
	fmt.Println("Jai Ganeshay Namah!!")

	// slice := make([]byte, 3, 15)
	usr := "Prateek"
	slice := []byte(usr)
	// We can also take a normal s lice of bytes and create a string from it with the simple conversion:
	str := string(slice)
	fmt.Println(str)

	// indexing a string yields its bytes, not its characters: a string is just a bunch of bytes. That means that when we store a character value in a string
	for i, _ := range usr{
		fmt.Println(usr[i])
	}
	
}

/*
Strings are actually very simple: they are just read-only slices of bytes with a bit of extra syntactic support from the language.

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string.


Because they are read-only, there is no need for a capacity (you can’t grow them), but otherwise for most purposes you can treat them just like read-only slices of bytes.

We can also take a normal slice of bytes and create a string from it with the simple conversion:

The Go language defines the word rune as an alias for the type int32


*/

/*
code point:
A code point is a numerical value that represents a specific character in Unicode

Code points are typically represented in hexadecimal form, such as U+0041 for the uppercase letter 'A'.

Characters:
Characters are identified and referenced by their Unicode code points. For example, the character 'A' is represented by the code point U+0041.

Unicode assigns a unique code point to each character, symbol, or glyph from various scripts, including Latin, Greek, Cyrillic, Chinese, Arabic, and many others.

Runes:

Rune is a Type. It occupies 32bit and is meant to represent a Unicode CodePoint.
As an analogy the english characters set encoded in 'ASCII' has 128 code points. Thus is able to fit inside a byte (8bit).
But guess what. There are many other symbols invented by humans other than the 'abcde..' symbols. And there are so many that we need 32 bit to encode them.
In golang then a string is a sequence of bytes. However, since multiple bytes can represent a rune code-point.

In Go, a rune is a data type that represents a Unicode code point. It is an alias for the int32 data type.

Runes are used to work with individual characters in a string, allowing you to manipulate and process Unicode characters efficiently.

*/

/*

Unicode Codepoint:

Imagine you have a big book that contains every single letter, number, and symbol in the world. Every single thing you can write or draw is in that book. Each thing has a special number assigned to it so we can find it in the book.

Now, when you want to talk about a specific letter, number, or symbol, you just say its special number, and everyone knows exactly what you mean. That special number is called a "Unicode codepoint."



Refer for more: https://go.dev/blog/strings

*/