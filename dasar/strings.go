package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aditya Pratama", "Adi"))
	fmt.Println(strings.Contains("Aditya Pratama", "Dodo"))
	fmt.Println(strings.ToLower("Aditya Pratama"))
	fmt.Println(strings.ToUpper("aditya pratama"))
	fmt.Println(strings.Trim("  Test  ", " "))
	fmt.Println(strings.ReplaceAll("Aditya Pratama Adit", "Adit", "Boba"))
	fmt.Println(strings.Split("Aditya Pratama", " ")) // split menjadi slice
}
