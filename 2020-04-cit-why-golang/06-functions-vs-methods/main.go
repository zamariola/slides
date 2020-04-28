package main

import "fmt"

type Person struct {
	Name string
}

func main() {

	person := Person{
		Name: "Leonardo",
	}
	//Function
	fmt.Println(FormatPersonName(person))
	//Method
	fmt.Println(person.FormattedName())

}

//Function
func FormatPersonName(p Person) string {
	return "Function: " + p.Name
}

//Method
func (p Person) FormattedName() string {
	return "Method: " + p.Name
}
