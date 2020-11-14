// Massive script on the entire overview of golang and its syntax as well as additional features
// Ryan McVicker
// must create a go.mod file when importing other packages
/*

	when using external packages:

		| - import the name of the package
		| - create a go.mod file with : go mod init "main.go"

*/
package main

import (
	"fmt"
	"rsc.io/quote"
)

func LoopsAndCons(){
	var loopName = "ryan mcvicker"
	//looping over a string or array etc
	//convert to string with string(value)
	for key, value := range loopName{
		fmt.Println("| index : ",key,"| value : ", string(value))
	}

	//similar to c for loops
	sum := 0
	for i := 0; i < 10; i++{
		sum += i
	}
	fmt.Println("Sum : ", sum)
}

func main() {
	//fmt is the main package used to print lines
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	LoopsAndCons()
}
