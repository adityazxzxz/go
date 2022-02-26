package main

// gunakan blank identifier (_) bila membuat sesuatu variabel atau import package tp g dipakai
import (
	"belajar-golang-dasar/database"
	// _ "belajar-golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
