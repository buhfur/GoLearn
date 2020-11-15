
package main


import (

	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)


func main(){



	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}
//handler function will be used to display the webpage
func index(w http.ResponseWriter, r *http.Request){
	

	
	//insert index.html here
	stream, err := ioutil.ReadFile("static/index.html")

	if err != nil { log.Fatal(err) }

	indexFile := string(stream)

	fmt.Fprintf(w, indexFile)

	//should this be in the main func?

//	defer func() { ioutil.NopCloser(stream) }()

}
