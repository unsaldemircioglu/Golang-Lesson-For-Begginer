package main

import "fmt"

type Student struct {
	name     string
	number   int
	lectures []string
}

func main() {
	s1 := Student{
		name:     "Unsal",
		number:   123,
		lectures: []string{"Matematik", "Kimya"},
	}
	// s2 := Student{
	// 	name:     "Unsal2",
	// 	number:   234,
	// 	lectures: []string{"Matematik", "Edebiyat"},
	// }

	fmt.Println("Ogrencinin İsmi: ", s1.name, "\n")
	fmt.Println("Ogrencinin Numarası: ", s1.number, "\n")
	fmt.Println("Aldigi Dersler: ", s1.lectures, "\n")

	// fmt.Println(s2)
}
