package main

import "fmt"

func NewName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var name map[string]string = NewName("Adit")
	if name == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(name["name"])
	}

}
