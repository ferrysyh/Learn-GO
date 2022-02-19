package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var number = []int{3, 7, 9, 10, 8, 6, 6, 3, 4, 5, 2}
	fmt.Println("Number :", number)

	var channell = make(chan float64)
	go getratarata(number, channell)
	var channel2 = make(chan int)
	go nilaimaksimal(number, channel2)

	for i := 0; i < 2; i++ {
		select {
		case rata_rata := <-channell:
			fmt.Printf("Rata-rata \t : %.2f \n", rata_rata)
		case maksimal := <-channel2:
			fmt.Printf("Maksimal \t : %d \n", maksimal)
		}
	}
}

func getratarata(number []int, ch chan float64) {
	var sum = 0
	for _, e := range number {
		sum += e
	}
	ch <- float64(sum) / float64(len(number))
}

func nilaimaksimal(number []int, ch chan int) {
	var max = number[0]
	for _, e := range number {
		if max < e {
			max = e
		}
	}
	ch <- max
}
