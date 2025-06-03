package main

import "fmt"

type Person struct {
	Name     string
	Age      uint8
	Document string
	Address  string
}

func main() {
	p := Person{
		Name:     "Roberto Filho",
		Age:      24,
		Document: "000.000.000-00",
		Address:  "Rua Teste",
	}

	fmt.Printf("Name: %s\nAge: %d\nDocument: %s\nAddress: %s", p.Name, p.Age, p.Document, p.Address)
}
