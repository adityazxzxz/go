package main

import "fmt"

type Person struct {
	Nama string
}

func (person *Person) sayName() {
	person.Nama = "Mr. " + person.Nama
}

func main() {
	nama := Person{
		Nama: "Adit",
	}

	nama.sayName()
	fmt.Println(nama)
}
