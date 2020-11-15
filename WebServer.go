//web server in go
 

package main


import (
	"fmt"
	"net/http"
)


func main(){



	//create handler to create website endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	// Determines which port the webserver will listen on	
	http.ListenAndServe(":8080", nil)

}
//handler function will be used to display the webpage
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello world</h1>")

}



func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello earth\b")
}
