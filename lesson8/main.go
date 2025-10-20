package main

import "fmt"

func changeName(m string) {
	m = "Unsal"
}

func main() {
	name := "Unsal1"
	p := &name

	fmt.Println("name degiskeni adresi:", &name)
	fmt.Println(p)  // 0xc000026070
	fmt.Println(*p) // bu adresteki değieri getir response Unsal1

	*p = "Unsal2"
	fmt.Println(name)

}
