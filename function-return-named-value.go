package main

import "fmt"

func getFullName()(firstName string, lastName string){
	firstName = "Aditya"
	lastName = "Pratama"

	return // tidak perlu return variable, karena sudah ada di atas sesudah parameter
}

func main(){
	a,b := getFullName()
	
	fmt.Println(a,b)
}