// panic digunakan untuk menghentikan program go
// apabila didalam function terdapat defer, maka defer tersebut tetap dijalankan walau panic
// fungsi recover adalah untuk menangkap message panic
// jika terdapat recover, maka program yg sudah di panic tidak akan berhenti
// agar recover berfungsi, maka harus diletakan didalam defer karena defer akan tetap berjalan walau panic, dengan begitu recover tetap bisa dijalankan karena berada didalam defer



package main

import "fmt"

func endapp(){
	panic_message := recover()
	fmt.Println("Panic message =",panic_message)
	fmt.Println("Aplikasi Selesai")
}

func firstapp(error bool){
	defer endapp()
	if error {
		panic("Aplikasi ERROR")
	}

	fmt.Println("Aplikasi Normal")
}

func main(){
	firstapp(false)
}