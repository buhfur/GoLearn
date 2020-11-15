package main

import (

	"fmt"
	"strings"
	"sort"
	"os"
	"log"
	"io/ioutil"
)

// common string functions
// file input and output




func main(){

	//file io in go
	file, err := os.Create("samp.txt")
	// if an error occurs 	
	if err != nil {

		log.Fatal(err)

	}

	file.WriteString("this is ome random text")

	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")
	
	if err != nil {

		log.Fatal(err)

	}

	readString := string(stream)
	fmt.Println(readString)
}





func stringFunctions(){


	sampString := "hello world"
	//check if string contains another string or pieces of a string
	fmt.Println(strings.Contains(sampString, "lo"))

	//index a part of a string
	fmt.Println(strings.Index(sampString,"lo"))
	// count the occurences of something in a string
	fmt.Println(strings.Count(sampString,"l"))

	//can also replace certain values
	//replace every l with x for the first 3 occurences
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	//can also split strings on a certain character
	fmt.Println(strings.Split(sampString,"l"))

	listOfletters := []string{"b","a","c"}

	//sort the list in alphabetical order
	sort.Strings(listOfletters)
	
	fmt.Println("letters : " , listOfletters)




	
}

