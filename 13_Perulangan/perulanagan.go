package main

import "fmt"

func main() {
	// cara 1 == for
	for i := 0; i < 10; i++ {
		fmt.Println("angka", i)
	}

	//cara 2 == while
	var j = 0
	for j < 10 {
		fmt.Println("angka", j)
		j++
	}

	// cara 3 == for each
	var k = 0
	for {
		fmt.Println("angka", k)
		k++
		if k == 10 {
			break
		}
	}
}
