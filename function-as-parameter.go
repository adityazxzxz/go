package main

import "fmt"

// type AliasFilter func(string) string

func filteredName(name string, filter func(string) string){
	fmt.Println(filter(name))
}

func checker(name string) string {
	if name == "anjing"{
		return "..."
	}else{
		return name
	}
}

func main(){
	filteredName("anjing",checker)
	filteredName("adit",checker)
}