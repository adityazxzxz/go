// Dalam interface terdapat kontrak yang terhubung dengan method GetName
// Secara otomatis apa bila ada function bernama GetName, maka method tersebut adalah bagian dari interface HasName
// method GetName dibuat dengan struct function

package main
import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}


func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasName HasName){
	fmt.Println("Hallo",hasName.GetName())
}

func main(){
	adit := Person{
		Name:"Adit",
	}
	SayHello(adit)

	tiger := Animal{
		Name:"Tiger",
	}
	SayHello(tiger)
}