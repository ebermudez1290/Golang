package main

import "fmt"

func main(){
	genericContactInfo:=contactInfo{ 
		email:"bot@gmail.com" ,
		zipCode:77339,
	}

	//Pointer init.
	Erick:= new(person)
	Erick.firstName = "Erick"
	Erick.lastName = "Bermudez"
	Erick.contactInfo = genericContactInfo
	fmt.Println(*Erick)

	//Auto properties
	Vero:= person{"Veronica","Marin",genericContactInfo }
	Vero.toString()

	//Named properties init
	Kino:= person{firstName:"Kino", lastName: "Bermudez", contactInfo:genericContactInfo}
	Kino.toString()

	//This definition assigns 0 values to the properties.
	var alex person
	alex.firstName = "Alex"
	alex.toString()


	jim:=person{
		firstName:"Jim", 
		lastName:"Jameson", 
		contactInfo: contactInfo{ 
			email:"jimjameson@gmail.com" ,
			zipCode:77339,
		},
	}

	//This will not be updated in the struct definition.
	jim.updateName("Jimmy")
	jim.toString()

	//This is the syntax that will use points vars to send and update reference values.
	jimPointer:= &jim
	jimPointer.updateLastName("Marin")
	jim.toString()
	(*jimPointer).toString()

	//GO Shortcut for pointers. it will take a value and behind the scenes turn it into a pointer to match the signature.
	jim.updateLastName("Mora")
	jim.toString()
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}


func (p person) toString(){
	fmt.Printf("%+v",p)
}

//Receiver that updates the first name BUT RECEIVES A PERSONAS AS VALUE TYPE.
func (p person) updateName(newFirstName string){
	p.firstName=newFirstName
}

//Receiver that updates the last name BUT RECEIVES A PERSONAS AS VALUE TYPE.
func (pointerToPerson *person) updateLastName(newLastName string){
	(pointerToPerson).lastName=newLastName
}