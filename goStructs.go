package main

import (
	"fmt"
)

func main(){


	rect1 := Rectangle{0,50,10,10}
	fmt.Println("rectangle is : ", rect1.width, " wide")

	fmt.Println("rectangle is : ", rect1.height, " tall ")

	fmt.Println("rectangle area is : ", rect1.area(), " cm^2")
	
	firstPerson := Person{"ryan mcvicker", 19, true}
	secondPerson := Person{"aiden pierce", 45, false}
		
	firstPerson.compare(secondPerson)
	needsPersonObject(firstPerson)	

}

type Person struct{
	name string
	age int
	isCitizen bool
}

type Rectangle struct{

	leftX float64
	topY float64
	height float64
	width float64
}
//define method for struct
func (rect *Rectangle) area() float64{
	
	return rect.width * rect.height

}

// can you make a struct function take another object as an argument
func (person *Person) compare(person2 Person) string{
	fmt.Println("Person 1 name : ", person.name, " | Person 2 name : ",person2.name)
	fmt.Println("Person 1 age : ", person.age, " | Person 2 age : ",person2.age)
	fmt.Println("Person 1 isCitizen : ", person.isCitizen, " | Person 2 isCitizen : ",person2.isCitizen)
	return "Compare Completed"
}

func needsPersonObject(randPerson Person){
	fmt.Println("this function has a person whos name is : ", randPerson.name)
}




