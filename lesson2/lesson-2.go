package main

import "fmt"

func main() {
	age := 13

	// fmt.Println(age > 10)
	// fmt.Println(age >= 30)
	// fmt.Println(age == 30)
	// fmt.Println(age != 30)
	/*
		true
		true
		true
		false
	*/

	if age > 18 {
		fmt.Println("Yasiniz 18'den buyuk, iceriye girebilirsiniz")
	} else if age >= 12 {
		fmt.Println("Yasiniz 12'ye esit veya buyuktur")
	} else if age >= 10 {
		fmt.Println("Yasiniz 10'ye esit veya buyuktur")
	} else {
		fmt.Println("Yasiniz cok kucuk")
	}

}
