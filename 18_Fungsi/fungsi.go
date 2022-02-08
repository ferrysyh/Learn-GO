package main

import "fmt"

func main() {
	tampilkan_pesan()
	fmt.Println(mengembalikan_nilaistr())
	fmt.Println(mengembalikan_nilaiint())
	fmt.Println(mengirim_nilai(10, 5))

}

// fungsi menampilkan pesan
func tampilkan_pesan() {
	fmt.Println("pesan berhasil diterima")
}

// fungsi mengembalikan nilai berupa string
func mengembalikan_nilaistr() string {
	return "pesan berhasil diterima"
}

// fungsi mengembalikan nilai berupa integer
func mengembalikan_nilaiint() int {
	return 10
}

// fungsi mengirim nilai argumen/parameter
func mengirim_nilai(x int, y int) int {
	var z = x * y
	return z
}
