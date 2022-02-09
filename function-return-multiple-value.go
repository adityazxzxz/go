package main

import "fmt"

func getName()(string, string, string){
	return "Aditya","Pratama","zxzxz"
}

func main(){
	firstName,middleName,_ := getName() // karena di go tidak boleh ada variabel yg nganggur, jika tidak mau ambil return lastname dari parameter, bisa gunakan underscore
	fmt.Println(firstName,middleName)
}