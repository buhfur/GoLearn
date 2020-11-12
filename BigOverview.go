package main
//when using 3rd party packages remember to make a go.mod file with : go mod init "main.go"

import (
	"fmt"
	"math/rand"
	"math"
	"rsc.io/quote"
)
//can make list of variables with var keyword

// need to create go.mod file before importing any external lib


func add(x int, y int) int{
	return x + y
}

//if the types are the same you can write them differently like this
func otheradd(x,y int) int{
	return x + y
}

// leaving a return blank will cause a "naked return" and will return the two present named return values (x, y int)
func split(sum int)(x, y int){
	x = sum * 4/9
	y = sum - x
	return //naked return here, only use for short functions as it hurts readability
}l

func secondsplit(name string)(returnname string){
	returnname = "ryan"
	//function should just return "returnn

	return //naked return here, only use for short functions as it hurts readability
}

func ConsAndLoops() {

	//for loop in go

	//C style for loop
	var (
		name = "Ryan McVicker"
		age = 19
		isSingle = 1
	)

	//looping over strings, arrays etc
	for key, value := range name{
		fmt.Println(name[key])
	}
}

}

func main(){
	fmt.Println(quote.Go())
	//variables
	var x,y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)//convert f to unsigned int
	//unsigned int cannot be negative
	//signed ints can be negative but has a lower positve range
	fmt.Println(x, y, z)
	fmt.Println("my favorite number is : ", rand.Intn(10))
	//name is exported if it begins with a capital letter
	fmt.Println(math.Pi)

	var useradd = add(3, 5)//saving function out put to variable

	fmt.Println(useradd)
	fmt.Println(split(17))
	fmt.Println(secondsplit("john"))
	ConsAndLoops()
}
