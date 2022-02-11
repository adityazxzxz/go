package main

import "fmt"

type Customer struct {
	Nama string
	Alamat string
	Umur int
}

func main(){
	var dude Customer
	dude.Nama = "Dude"

	fmt.Println(dude)
	
	adit := Customer{
		Nama: "Adit",
		Alamat:"Tangerang",
		Umur: 22,
	}

	fmt.Println(adit)
}