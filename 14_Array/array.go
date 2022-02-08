package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "anggur", "melon"}
	buah[1] = "mangga" // mengubah value
	fmt.Println("Jumlah Element :", len(buah))
	fmt.Println("Isi Element :", buah)
}
