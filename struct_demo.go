package main

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
