package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	go tampilkan_pesan(5, "saya kirim")
	tampilkan_pesan(5, "saya kedua")

	// blocking untuk menahan karakter agar tidak mengakhiri sebuah proses
	var input string
	fmt.Scanln(&input)
}

func tampilkan_pesan(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
