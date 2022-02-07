package main

import "fmt"

func main() {
	var angka = 6

	// menggunakan if else
	if angka%2 == 0 {
		fmt.Println("bilangan genap")
	} else {
		fmt.Println("bilangan ganjil")
	}

	// menggunakan switch
	switch angka % 2 {
	case 0:
		fmt.Println("bilangan genap")
	default:
		fmt.Println("bilangan ganjil")
	}
}
