package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Println("Halo",firstName,lastName)
}

func main(){
	firstName := "Aditya"
	lastName := "Pratama"

	sayHello(firstName,lastName)
}