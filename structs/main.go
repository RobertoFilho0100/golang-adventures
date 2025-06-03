package main

import "fmt"

type Person struct {
	Name     string
	Age      uint8
	Document string
	Address  string
}

func main() {
	/*
	   Struct initialization can be done in two ways:

	   1. Using named fields (recommended for better readability and maintenance):
	      This approach clearly associates each value with its field.

	   2. Using positional values without field names:
	      The values must be provided in the exact order the fields are declared in the struct.
	      This method is shorter but can be error-prone if the order changes.

	   Example of positional initialization (not used here):
	       p2 := Person{"Roberto Filho", 24, "000.000.000-00", "Rua Teste"}
	*/

	p := Person{
		Name:     "Roberto Filho",
		Age:      24,
		Document: "000.000.000-00",
		Address:  "Rua Teste",
	}

	fmt.Printf("Name: %s\nAge: %d\nDocument: %s\nAddress: %s", p.Name, p.Age, p.Document, p.Address)
}
