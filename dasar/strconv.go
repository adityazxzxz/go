// strconv = string convertion
package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(angka)
	} else {
		fmt.Println(err.Error())
	}

	angka2, error := strconv.Atoi("2000")
	if error == nil {
		fmt.Println(angka2)
	} else {
		fmt.Println(error.Error())
	}

	var angka3 string = strconv.Itoa(123)
	if err == nil {
		fmt.Println(angka3)
	} else {
		fmt.Println(err.Error())
	}

	benar, _ := strconv.ParseBool("true")
	fmt.Println(benar)
}
