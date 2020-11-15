package main

//entire file dedicated to pointers in go
//cause they are fucking annoying to deal with in other languages

import "fmt"

func main(){



	x := 0
	//generate new pointer 
	yPtr := new(int)
	strPtr := new(string)	
	fmt.Println("Simple print out yPtr : " , *yPtr)
	fmt.Println("Simple print out strPtr : " , *strPtr)

	changeYVal(yPtr)
	changeStrPtr(strPtr)

	
	fmt.Println("yPtr after changing : " , *yPtr)
	fmt.Println("strPtr after changing : " , *strPtr)
	
	fmt.Println("value of x: " , x)
	// pass the memory address to the function
	changeXVal(&x)
	fmt.Println("value of x: " , x)
}


func changeXVal(x *int){
//	x = 2

	fmt.Println("var x before changing : ", x)
	fmt.Println("var x  addr : ", &x) // x *int = &x ; fmt.Println
	*x = 2 // store "2" at memory address of variable
}


func changeYVal(yPtr *int){

	*yPtr = 100 // store "2" at memory address of variable



}




func changeStrPtr(strPtr *string){

	*strPtr = "changed string pointer"



}

