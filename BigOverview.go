package main
//when using 3rd party packages remember to make a go.mod file with : go mod init "main.go"
//fmt is a format package
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
	//can make constant varibles with the const keyword
	const pi float64 = 3.14159265 
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
	ConsAndLoops()
	StringFormatting()
	Arrays()
	SwitchCases()
	afterMain()
	makingMaps()

}
//can declare functions after the main function

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

