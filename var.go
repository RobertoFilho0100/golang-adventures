package main

import "fmt"

/*
There are two common ways to declare a variable with a specific type and value in Go:
1. Explicit type declaration:
   var age uint8 = 24
2. Type inference with type conversion (casting):
   var age = uint8(24)
Both result in `age` being of type uint8 with value 24.
*/

var (
	age      = uint8(24)
	document = "000.000.000-00"
)

func main() {
	name := "Roberto Filho"
	address := "Rua Teste"

	fmt.Printf("Name: %s\nAge: %d\nDocument: %s\nAddress: %s\n", name, age, document, address)
}
