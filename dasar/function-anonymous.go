package main

import "fmt"

//type CheckName func(string) bool

func checkin(name string,checkname func(string) bool){
	if checkname(name){
		fmt.Println("you are blocked",name)
	}else{
		fmt.Println("Welcome",name)
	}
}

func main(){
	blacklist := func(name string) bool {
		return name == "Adit"
	}
	checkin("Adit",blacklist)
	checkin("Aditya",blacklist)
	checkin("Asu",func(name string) bool {
		return name == "Adit"
	})
}