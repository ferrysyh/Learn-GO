package main

import "fmt"

func main() {
	var mahasiswa = map[string]int{"Aldo": 182, "Yosep": 178}
	tampil_mahasiswa(mahasiswa)
}

func tampil_mahasiswa(mahasiswa map[string]int) {
	for nama, tinggi := range mahasiswa {
		fmt.Println(nama, " : ", tinggi, " cm")
	}
}
