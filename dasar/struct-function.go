package main

import "fmt"

type Customer struct {
	Nama string
	Alamat string
	Umur int
}

func (customer Customer) sayHi(){
	fmt.Println("Hi",customer.Nama)
}

func (customer Customer) sayBalik(nama string){
	fmt.Println("Hi",customer.Nama,", Saya",nama)
}

func main(){
	adit := Customer{
		Nama : "Adit",
		Alamat: "Tangerang",
	}

	adit.sayHi()
	adit.sayBalik("Dude")

}