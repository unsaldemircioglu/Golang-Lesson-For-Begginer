package main

import "fmt"

// func sayHello(name string) {
// 	fmt.Println("Hello", name)
// }
func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func main() {

	area1 := calcArea(5)
	fmt.Println(area1)

	// sayHello("Unsal")
	// sayHello("Sena")
}
