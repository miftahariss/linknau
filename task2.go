// interface
// interface adalah tipe data yang abstract, dia tidak memiliki implementasi secara langsung
// sebuah interface memiliki definisi definisi method
// dan biasanya interface digunakan sebagai kontrak, jadi kita biasanya bikin interface sebagai kontrak (harus ada yang mengimplementasikannya)
// biasanya interface itu diimplementasikan dalam bentuk sebuah struct
// karena interface itu sebuah kontrak jadi nanti implementasi struct nya itu boleh banyak, jadi bisa satu, dua dan seterusnya yang penting mengikuti kontrak yang sudah ditetapkan oleh si interface nya.
// Dalam contoh ini, struct Person mengimplementasikan interface Greeter karena menyediakan metode Greet().

package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var greeter Greeter = Person{Name: "John"}
	fmt.Println(greeter.Greet())
}
