package main

import "fmt"

func main() {
	var name string = "Unsal"
	var age int = 20
	var price float32 = 55.87

	// println
	fmt.Println("Selam \n")
	fmt.Println("Mrb", name, "yasiniz: ", age)
	//printf
	fmt.Printf("Adiniz: %v, Yasiniz %v \n", name, age)
	fmt.Printf("Adiniz: %q, Yasiniz %q \n", name, age)
	fmt.Printf("Degiskenimin tipi %T \n", name)
	fmt.Printf("Bu urunun fiyati: %f0.1f \n", price)

	//
	var myText string = fmt.Sprintf("Bu urunun ffiyatı: %0.2f \n", price)
	fmt.Println(myText)
}
