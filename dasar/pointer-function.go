package main

import "fmt"

type Users struct {
	Nama   string
	Alamat string
	Umur   int
}

func ubahStruct(user *Users) {
	user.Alamat = "Tangerang"
}

func main() {
	var user Users = Users{
		Nama:   "Adit",
		Alamat: "Karawaci",
		Umur:   28,
	}

	// ubahStruct(user)
	fmt.Println(user)
	ubahStruct(&user)
	fmt.Println(user)
}
