// interface kosong bebas me return value tanpa mementingkan type data
package main

import "fmt"

type kosong interface{}

func check(i int) kosong {
	if i == 1 {
		return 1 // return integer
	} else if i == 2 {
		return "dua" // return string
	} else {
		return false // return boolean
	}
}

func main() {

	// walaupun bisa menyesuaikan type data, saat declare variabel harus tetap bertype interface
	var result interface{}
	result = check(1)

	fmt.Println(result)
	fmt.Println(check(1))
	fmt.Println(check(2))
	fmt.Println(check(3))
}
