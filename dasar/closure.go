// closure mengakses suatu variable yg masih dalam 1 scope
// misalnya scope increment() function bisa mengakses variabel counter yg ada diluar blocknya
// akan tetapi, scope main() tidak dapat mengakses variabel total dalam func increment

package main

import "fmt"

func main(){
	counter := 0

	increment := func(){
		total := 10
		counter++
		fmt.Println(counter)
		fmt.Println(total)
	}

	increment()
}