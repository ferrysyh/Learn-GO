package main

import "fmt"

func main() {
	var nomor int = 10
	var tas string = "hello"
	var alamat_nomor *int = &nomor
	var alamat_tas *string = &tas

	fmt.Println("Nilai dari nomor : ", nomor)
	fmt.Println("Alamat variabel nomor : ", alamat_nomor)
	fmt.Println("Nilai dari nomor : ", tas)
	fmt.Println("Alamat variabel nomor : ", alamat_tas)
}
