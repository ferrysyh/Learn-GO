package main

import "fmt"

func main() {
	var nilai = 5

	// menggunakan if else
	if nilai == 10 {
		fmt.Println("lulus dengan sempurna")
	} else if nilai >= 7 {
		fmt.Println("lulus")
	} else if nilai == 6 {
		fmt.Println("sedikit lagi")
	} else {
		fmt.Println("belajar lagi")
	}

	// menggunakan switch
	switch nilai {
	case 10:
		fmt.Println("sempurna")
	case 9:
		fmt.Println("bagus")
	case 8:
		fmt.Println("lumayan")
	default:
		fmt.Println("belum lulus")
	}
}
