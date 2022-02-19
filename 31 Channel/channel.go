package main

import "fmt"
import "runtime"

var pesan = make(chan string)

func main() {
	runtime.GOMAXPROCS(2)

	go hello("a")
	go hello("b")
	go hello("c")

	var message1 = <-pesan
	fmt.Println(message1)

	var message2 = <-pesan
	fmt.Println(message2)

	var message3 = <-pesan
	fmt.Println(message3)
}

func hello(nama string) {
	var data = fmt.Sprintf("Hello %s", nama)
	pesan <- data
}
