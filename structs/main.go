package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// creating a new person struct
	//alex := person{"Alex", "Anderson"}

	// this syntax makes it easier to change order of fields are add more fields on. the tutorial teacher prefers this method.
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	//third way to instantiate a struct
	//var alex person

	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//alex.contact.email = "alex.anderson@gmail.com"
	//alex.contact.zip = 68130

	//fmt.Println(alex)
	// prints field names and values
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName: "Johnson",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 68130,
		},
	}

	
	// &variable - gives memory address of the value this variable is pointing at
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	// this accomplishes the same as above. Go allows shortcut of type of person into pointer person for you.
	jim.updateName("Jimmy")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person = type description - means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson = gives the value of this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}