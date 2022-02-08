package library

import "fmt"

type nama_mahasiswa struct {
	nama string
	umur int
}

// Public
func Public() {
	var m1 nama_mahasiswa
	m1.nama = "Aldo"
	m1.umur = 19
	var m2 = nama_mahasiswa{"John Mayer", 26}
	m1.tampil()
	m2.tampil()
}

// Private
func private() {
	var m1 nama_mahasiswa
	m1.nama = "Aldo"
	m1.umur = 19
	var m2 = nama_mahasiswa{"John Mayer", 26}
	m1.tampil()
	m2.tampil()
}

func Public2() {
	private()
}

func (s nama_mahasiswa) tampil() {
	fmt.Println("Nama Saya Adalah :", s.nama)
	fmt.Println("Nama Saya Adalah :", s.umur)
	fmt.Println("")
}
