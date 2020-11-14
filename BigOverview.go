package main
//Ryan McVicker
//when using 3rd party packages remember to make a go.mod file with : go mod init "main.go"
//fmt is a format package
import (
	"fmt"
	//"math/rand"
	//"math"
	//"rsc.io/quote"
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
}

func ConsAndLoops() {

	//for loop in go

	//C style for loop
	var (
		name = "Ryan McVicker"
		age = 19
		isSingle = 1
	)
	// using "range" for parsing through a string or array etc
	//looping over strings, arrays etc
	for key, value := range name{
		fmt.Println("Index : ", key, "| value : ",string(value))
	}
	
	for i := 0; i < age; i++{
		
		fmt.Println("Ryan McVicker is : ",age)
	}
	fmt.Println("is ryan single? : ", isSingle)
		
}

func StringFormatting(){
	//can format strings with fmt.Printf()
	var (
		firstName = "ryan"
		lastName = "mcvicker"
		age = 19
		)
		// use %d for formatting a decimal and not %i
		fmt.Printf(" Users fist, last name, and age : %s\n %s\n %d\n",firstName,lastName,age)
	}
func SwitchCases(){
	yourAge := 18
	
	switch yourAge {
		case 16: fmt.Println("go drive")
		
		case 18: fmt.Println("go vote")
		case 21: fmt.Println("go drink alcohol")
		default : fmt.Println("go have fun") 
	}	

}



func Arrays(){
	var favNums2[5] float64
	favNums2[0] = 263
	favNums2[1] = 363
	favNums2[2] = 463
	favNums2[3] = 563
	
	fmt.Println(favNums2[3])
	
	//can also initialize an array like this

	//create arrray without specified length
	randArray := []int {6,7,8,9,10,11}
	
	for key,value := range randArray{
		fmt.Println("random array : " ,key,value)
	// count from 1 instead of zero when creating array in brackets
	}
	favNums3 := [5]int {1,2,3,4,5}
	//call AddUp on array here
	fmt.Println("sum of array is : ",AddUp(randArray))

	favStrings := [5]string{"ryan","bob","joe","lisa","bart"}
	for key,value := range favNums3	{
		fmt.Println(key,value)
	}
	//array of strings
	for key,value := range favStrings{
		fmt.Println(key,string(value))
	}
}

func main(){
	//initialize variables here
	//const pi float64 = 3.14159265 

	//var x,y int = 3, 4

	//var f float64 = math.Sqrt(float64(x*x + y*y))

	//var z uint = uint(f)//convert f to unsigned int

	num1, num2 := addTwoValues(5)

	//make closures to access variables in main in another function
	num3 := 3

	doubleNum := func() int {
		num3 *= 2

		return num3
	}

	//call functions here
	fmt.Println("variaticFunction Output : ",variaticFunction(1,2,3,4,5))
	fmt.Println("num1 : ", num1, "\n Num2 : ", num2)
	fmt.Println("closure output : " , doubleNum())
	fmt.Println("closure output 2 : " , doubleNum())
	//will fail to test exception
	fmt.Println(safeDiv(3,0))
	//calls the chosen defer function after main() is executed
	// used to call cleanup functions like closing a file after opening
	defer printTwo() 
	printOne()

	demPanic()	
}

func afterMain(){
	fmt.Println("this is after the main function")

}

func makingMaps(){
	//function about maps in go (or dictionaries as i call them)
	presAge := make(map[string] int)
	presAge["George Washington"] = 64
	// easilly print out the length of an array 
	fmt.Println(len(presAge))
	fmt.Println(presAge["George Washington"])
	
	fmt.Println(presAge)
}

func AddUp ( number []int ) int{
	sum := 0

	for _, val := range number{
		sum += val
	}
	return sum
}



// can make function take 1 input and return x amount of values
func addTwoValues(num int) (int, int){

	return num+1, num+2 

}


func variaticFunction(args ...int) int {
	//function with undefined amount of arguments
	//guessing that these would be used for command line interfaces???
	finalValue := 0 
	for _, value := range args{
		finalValue -= value
	}
	return finalValue
}

//recursive functions 

func factorial(num int) int{

	if num == 0 {
		return 1
	}
	
	return num * factorial(num - 1)
}


func printOne() { fmt.Println(1) } 
func printTwo() { fmt.Println(2) } 



func safeDiv(num1, num2 int) int {

	defer func(){
		//will catch any exceptions 
		//recover allows the program to continue execution after encountering an error
		fmt.Println(recover())

	}()
	solution := num1 / num2
	return solution
}


func demPanic(){

	//used to define error handling exceptions
	defer func(){
		fmt.Println(recover())
	}()

	panic("PANIC")
}
