package main

import (
	"fmt"
	"os"
)

func main() {
	// Dosya yoksa oluştur, varsa sonuna ekle
	file, err := os.OpenFile("ornek.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Yeni satır eklendi!\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("Dosyaya veri eklendi.")
}
