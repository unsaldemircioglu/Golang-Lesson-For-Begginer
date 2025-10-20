package student

import "fmt"

type Student struct {
	name     string
	number   int
	lectures []string
}

func Create(name string, number int, lectures []string) Student {
	return Student{
		name:     name,
		number:   number,
		lectures: lectures,
	}
}

func ShowInfo(s Student) {
	fmt.Println("Ogrencinin İsmi: ", s.name, "\n")
	fmt.Println("Ogrencinin Numarası: ", s.number, "\n")
	fmt.Println("Aldigi Dersler: ", s.lectures, "\n")
}
