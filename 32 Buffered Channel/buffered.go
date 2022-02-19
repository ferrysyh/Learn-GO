package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var pesan = make(chan int, 2)

	go func() {
		for {
			var i = <-pesan
			fmt.Println("terima data", i)
		}

	}()

	for i := 0; i < 5; i++ {
		fmt.Println("kirim data", i)
		pesan <- i
	}
}
