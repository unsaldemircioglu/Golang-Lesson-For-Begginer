package main

import "fmt"

func main() {
	//Array
	students := [3]string{"Unsal", "Unsal2", "Unsal3"}

	students[0] = "Unsal123"

	// Slice
	salaries := []int{100, 200, 120, 500}
	salaries = append(salaries, 500)
	salaries = salaries[0:2]

	fmt.Println(students[0])
	fmt.Println(salaries)
	fmt.Println(len(salaries))

	//Slice
}
