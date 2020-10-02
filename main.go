package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "John"
	fmt.Println(*firstName)

	lastName := "Doe"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Smith"
	fmt.Println(ptr, *ptr)

}
