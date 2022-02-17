// variabel orang2 adalah pointer
// artinya jika isi variabel orang2 di ubah, maka reference yaitu si orang1 maka akan ikutan berubah datanya
// mudahnya, orang2 nge-link ke orang1

package main

import "fmt"

type Person struct {
	Nama, Alamat string
	Umur         int
}

func main() {
	var orang1 Person = Person{"Adit", "Harkit", 23}

	// cara membacanya
	// var namavariabel typedatanya
	// dibawah berarti nama variabel nya orang2, type datanya struct tapi pointer karena ada * nya
	// dan &orang1 artinya mengarah/mereferensikan ke orang1
	var orang2 *Person = &orang1

	orang2.Nama = "Joko"

	fmt.Println(orang1)
	fmt.Println(orang2)

	//jadi orang2 adalah sebuah pointer. maka saat mencetak reference hanya memory address
	fmt.Println(&orang2)
	fmt.Println(*orang2)
}
