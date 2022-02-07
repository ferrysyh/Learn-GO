package main

import "fmt"

func main() {
	var buah = []string{"apel", "mangga", "jeruk", "anggur"}
	buah = append(buah, "durian")
	fmt.Println(buah)
	fmt.Println("Jumlah Element : ", len(buah))
}