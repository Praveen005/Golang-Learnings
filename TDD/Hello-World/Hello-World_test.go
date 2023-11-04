package main

import (
	"testing"
)

// func TestHello(t *testing.T){
// 	got := hello("Praveen")
// 	want := "Hello, Praveen"

// 	if got != want{
// 		t.Errorf("got %q want %q", got , want)
// 	}
// }

/*
func TestHello(t *testing.T){
	t.Run("Saying Hello to People", func(t *testing.T){
		got := hello("Praveen")
		want := "Hello, Praveen"

		if got != want{
			t.Errorf("got %q want %q", got , want)
		}
	})

	t.Run("Say, 'Hello, world' when no string is supplied", func(t *testing.T){
		got := hello("")
		want := "Hello, world"

		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	})
}

*/

/*
What have we done here?
	We've refactored our assertion into a new function. This reduces duplication and improves readability of our tests

For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy

t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output. 

*/
/*
func TestHello(t* testing.T){
	t.Run("Saying Hello to people", func(t * testing.T){
		got := hello("Praveen");
		want := "Hello, Prveen"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world', when empty string is passed", func(t *testing.T) {
		got := hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
*/

func TestHello(t *testing.T){
	t.Run("Saying Hello to the user in Spanish", func(t *testing.T) {
		got := hello("Praveen", "Spanish")
		want := "Hola, Praveen"
		assertCorrectMessage(t, got, want)
	})


	t.Run("Print 'Hello, world' when the string entered is empty", func(t *testing.T) {
		got := hello("", "Spanish")
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in French", func(t *testing.T) {
		got := hello("Praveen", "French")
		want := "Bonjour, Praveen"

		assertCorrectMessage(t, got, want)
	})
}


func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if(got != want){
		t.Errorf("got %q, want %q", got, want);
	}
}









/*

To run the test: go test


Writing a test is just like writing a function, with a few rules
	It needs to be in a file with a name like xxx_test.go
	The test function must start with the word Test(here: TestHello())
	The test function takes one argument only t *testing.T
	In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file


	 For tests %q is very useful as it wraps your values in double quotes.


	 t.Run:  This function allows you to create subtests within a test function. It's useful for organizing and grouping related tests

	 In 2nd TestHello() : we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.
*/