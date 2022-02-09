package main

import "fmt"

func main(){
	nama := "Adits"

	switch nama {
	case "Adit":
		fmt.Println("Ini adit")
	default:
		fmt.Println("ini bukan adit")	
	}

	switch {
	case len(nama) > 10:
		fmt.Println("Nama terlalu panjang")
	case len(nama) > 3:
		fmt.Println("Nama lumayan panjang")
	}


	switch length := len(nama); length > 3 {
	case true:
		fmt.Println("Nama lebih dari 3 huruf")
	case false:
		fmt.Print("Nama kurang dari 3 huruf")
	}

}