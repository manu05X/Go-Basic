package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

/*
	type person struct {
		firstName string
		lastName  string
		contact   contactInfo
	}
*/
type person struct {
	firstName string
	lastName  string
	contactInfo
}

/*
func main() {
	// alex := person{"Alex", "Andreson"} // 1st
	// manish := person{firstName: "Manish", lastName: "Kumar"} //2nd way

	// fmt.Println(alex, manish)

	// var alex person //3rd way
	// fmt.Println(alex) // { }
	// fmt.Printf("%+v", alex) // {firstName: lastName:}%

	var alex person //3rd way

	alex.firstName = "Manish"
	alex.lastName = "Kumar"
	fmt.Println(alex)       // {Manish Kumar}
	fmt.Printf("%+v", alex) // {firstName:Manish lastName:Kumar}%
}
*/

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "hdksk@gmail.com",
			zipCode: 785332,
		},
	}

	//fmt.Printf("%v", jim) //{Jim Party {hdksk@gmail.com 785332}}%
	//fmt.Printf("%+v", jim) //{firstName:Jim lastName:Party contactInfo:{email:hdksk@gmail.com zipCode:785332}}%

	//jim.print() //{firstName:Jim lastName:Party contactInfo:{email:hdksk@gmail.com zipCode:785332}}%

	// //jim.updateName("Manish")
	// jimPointer := &jim // &jim give me the memory address of the value this variable is ponting
	// jimPointer.updateName("Manish")
	// jim.print()

	//jim.updateName("Manish")
	jim.updateName("Manish")
	jim.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// use pointer to update the first name
// here *person is a pointer type and variable name is pointreToPerson
func (pointreToPerson *person) updateName(newFirstName string) {
	//*pointreToPerson -> whenever we place a * in front of a variable it will be converted into a value, i.e it is a operator and we want to manipulate the value the pointer is referencing.
	(*pointreToPerson).firstName = newFirstName //*pointreToPerson -> Give me the value this memory address is pointing at, so pointreToPerson is the memory address where jim exist at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/* go is a pass by value language means whenever we pass a vlue to a function go will take the value or that struct and copy
all the value/data inside that struct and place it inside some new container inside our computer memory
*/
