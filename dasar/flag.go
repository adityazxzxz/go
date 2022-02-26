// jalankan go run dasar/flag.go -user=adit
// jika tidak ditulis argumennya, maka muncul default valuenya yaitu root
package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("user", "root", "Masukan username")
	flag.Parse()

	fmt.Println("Nama : ", *username)
}
