package main

import "fmt"

func main() {

	x := 1
	// for x < 5 {
	// 	fmt.Println("x'in degeri : ", x)
	// 	// x = x + 1
	// 	x++
	// }

	//Infinity Loops(Sonsuz Döngüler)
	// for x > 0 {
	// 	fmt.Println("x'in degeri : ", x)
	// 	x++
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("x'in degeri: ", x)
	}

	students := []string{"Ahmet", "Mehmet", "Zeynep", "Unsal", "Sena"}

	for i := 0; i < len(students); i++ {
		// fmt.Println(i)
		fmt.Println(students[i])
	}

	for index, value := range students {
		fmt.Println(index)
		fmt.Println(value)
	}

}
