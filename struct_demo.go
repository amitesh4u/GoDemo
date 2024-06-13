package main

import "fmt"

type Person struct {
	firstname   string
	lastName    string
	ContactInfo // Here we can skip the field name i.e. contact ContactInfo
}

// Structs is pass by value so direct update won't work. We have to pass the memory address
func (pointerToPerson *Person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

type ContactInfo struct {
	mobileNo string
	email    string
}

func testStructPointerFlows() {
	person := Person{
		firstname: "fName",
		lastName:  "lName",
		ContactInfo: ContactInfo{
			mobileNo: "9876543210",
			email:    "test@test.com",
		},
	}

	fmt.Print(person) // Without Key
	//{fName lName {9876543210 test@test.com}}

	fmt.Printf("\n%+v", person) // With Key: Value
	//{firstname:fName lastName:lName ContactInfo:{mobileNo:9876543210 email:test@test.com}}

	personPointer := &person // Passing the memory address
	personPointer.updateFirstName("updatedFName")
	fmt.Printf("\n%+v", person)

	//OR in shortcut. Go will convert it to memory address if the receiver is a pointer
	person.updateFirstName("updatedAgainFName")
	fmt.Printf("\n%+v", person)
}
