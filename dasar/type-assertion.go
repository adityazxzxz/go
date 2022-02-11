package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println(value, "adalah string")
	case int:
		fmt.Println(value, "adalah int")
	default:
		fmt.Println("unknown")
	}

}
