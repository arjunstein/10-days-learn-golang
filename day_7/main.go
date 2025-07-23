package main

import "fmt"

// Interface Hewan
type Hewan interface {
	Suara()
}

// Struct Kucing
type Kucing struct {
	Nama string
}

func (k Kucing) Suara() {
	fmt.Println(k.Nama, "bilang: Meong!")
}

// Struct Anjing
type Anjing struct {
	Nama string
}

func (a Anjing) Suara() {
	fmt.Println(a.Nama, "bilang: Guk guk!")
}

// Struct Ayam
type Ayam struct {
	Nama string
}

func (ay Ayam) Suara() {
	fmt.Println(ay.Nama, "bilang: Kukuruyuk!")
}

// Fungsi polymorphic
func TampilkanSuara(h Hewan) {
	h.Suara()
}

func main() {
	// Buat slice berisi berbagai hewan
	hewanList := []Hewan{
		Kucing{Nama: "Kitty"},
		Anjing{Nama: "Bruno"},
		Ayam{Nama: "Jago"},
	}

	// Tampilkan semua suara hewan
	for _, h := range hewanList {
		TampilkanSuara(h)
	}
}
