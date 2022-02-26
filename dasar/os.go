package main

import (
	"fmt"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}
}
