//
// run this : go run 2.go
// Making a WebServer
//
// ex: http://localhost:9090/
// ex: http://localhost:9090/?url_long=111&url_long=222
//
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

//
func sayHelloWorld(w http.ResponseWriter, r *http.Request){
	r.ParseForm() 
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello Fantastic!") // send data to client
}
// Application Entry
func main(){
	fmt.Println("Go 102")
	http.HandleFunc("/",sayHelloWorld) // Set Router
	err := http.ListenAndServe(":9090", nil) // Set Port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} 
}