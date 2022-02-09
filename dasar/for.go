package main
import "fmt"

func main(){
	

	for counter:=0;counter <10; counter++{
		fmt.Println("perulangan ke",counter)
	}

	slice := []string{"Adit","Eko","Babo"}

	for i:=0;i<len(slice);i++{
		fmt.Println(slice[i])
	}

	for index,value := range slice{
		fmt.Println(index,value)
	}

	mahasiswa := make(map[string]string)
	mahasiswa["nim"] = "12345"
	mahasiswa["nama"] = "Adit"

	for key,value := range mahasiswa {
		fmt.Println(key,value)
	}
}