package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _,number := range numbers {
		total += number
	}

	return total
}

func main(){
	result := sumAll(12,12,12)

	fmt.Println(result)
}