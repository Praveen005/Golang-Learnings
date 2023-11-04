package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main(){

	//when a request comes in and matches this path, executes the function in 2nd parameter
	//HandleFunc is a function fron the net/http package, it registers a fuction to a path on a thing called DefaultServeMux(It is an httphandler and everything related to server in Go is an http handler)
	// DefaultServeMux is a variable in the Go that represents the default HTTP request multiplexer (or multiplexor) for the built-in HTTP server provided by the Go standard library. The multiplexer is responsible for routing incoming HTTP requests to the appropriate handlers based on the requested URL path.
	// http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
	// 	log.Println("Hello world!")
	// })

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		//reading using the http request
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "Oops", http.StatusBadRequest)  //another way of replacing writing below two lines, by inbuilt http.Error() method
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))
			return //it is still required, because error doesn't stop the flow of my application
		}
		log.Printf("Data %s\n", d)
		//writing back to the client, using response writer

		fmt.Fprintf(rw, "Hello %s\n", d)
	})


	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Good Bye, world!")
	})

	//If you want to explicitly bind it to some IP(local host, port 9090 here)
	// http.ListenAndServe("127.0.0.1:9090", nil)
	//if you want to bind to every ip
	http.ListenAndServe(":9090", nil) //2nd parameter here is the handler
}