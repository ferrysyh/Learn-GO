package main

import "fmt"

func main() {
	var x1 pelajar
	var x2 pelajar

	x1.nama = "johan"
	x1.kelas = 12
	x1.umur = 17

	x2.nama = "alvin"
	x2.kelas = 11
	x2.umur = 16

	fmt.Println("nama :", x1.nama)
	fmt.Println("kelas :", x1.kelas)
	fmt.Println("umur :", x1.umur)
	fmt.Println("===========")
	fmt.Println("nama :", x2.nama)
	fmt.Println("kelas :", x2.kelas)
	fmt.Println("umur :", x2.umur)
}

type pelajar struct {
	nama  string
	kelas int
	umur  int
}
