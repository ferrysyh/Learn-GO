package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")

	fmt.Println("Jumlah Element : ", len(buah))
	fmt.Println("Isi Element : ", buah)

	fmt.Println("=========== cara 1 ==========")
	for i := 0; i < len(buah); i++ {
		fmt.Println("Element ke - :", i, buah[i])
	}
	fmt.Println("=========== cara 2 ==========")
	var j = 0
	for j < len(buah) {
		fmt.Println("Element ke - :", j, buah[j])
		j++
	}
	fmt.Println("=========== cara 3 ==========")
	var k = 0
	for {
		fmt.Println("Element ke - :", k, buah[k])
		k++
		if k == len(buah) {
			break
		}
	}
}
