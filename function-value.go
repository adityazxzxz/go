package main

import "fmt"

func sayGoodBye(nama string) string {
	return "Good bye "+nama
}

func main(){
	goodBye := sayGoodBye // function saygoodby dimasukin ke variabel. jadi nanti variabel goodby bisa di kirim melalui parameter lagi. istilahnya passing function lewat parameter bisa pakai ini

	fmt.Println(goodBye("Dude"))
	fmt.Println(sayGoodBye("Dude"))
}