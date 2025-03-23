package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

type Employee struct {
	Person
	Position string
	Company  string
}

func main() {
	emp := Employee{
		Person: Person{
			Name:    "Adiva Nursuandy Ritonga",
			Age:     25,
			Address: "Medan",
		},
		Position: "BackEnd Developer",
		Company:  "Bimasakti Multi Sinergi",
	}

	fmt.Println(emp.Person.Name)
	fmt.Println(emp.Position)
	fmt.Println(emp.Company)
}
