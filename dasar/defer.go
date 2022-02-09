// defer memanggil function lain setelah function dieksekusi
// contoh dibawah apabila firstApp dipanggil dan selesai maka, endApp akan dijalankan 
// selalu letakan defer di awal function

package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi Selesai")
}

func firstApp(error bool){
	defer endApp()
	if error {
		fmt.Println("Error")
	}

	fmt.Println("Aplikasi Normal")
}

func main(){
	firstApp(true)
}