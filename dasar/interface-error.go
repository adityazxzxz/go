package main

import (
	"fmt"
	"errors"
)

func pembagian(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh kosong")
	}else{
		return angka / pembagi, nil
	}
}

func main(){
	hasil, err := pembagian(4,0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
	
}