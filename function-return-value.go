package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "Hello Bro"
	}else{
		return "Hello " + name
	}
}

func main(){
	result := sayHello("Adit")

	fmt.Println(result)
	fmt.Println(sayHello("Bagus"))
	fmt.Println(sayHello(""))
}