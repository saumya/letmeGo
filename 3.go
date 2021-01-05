//
// Run This : go run 3.go
// Making a file server
//
// ex: http://localhost:8000/
//

package main
//
import (
	"fmt"
	"net/http"
)
//
func ping(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("pong"))
}
//
// Application Entry
func main(){
	fmt.Println("Go 103 : File Server")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}