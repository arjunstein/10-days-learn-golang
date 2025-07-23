package main

import (
	"fmt"
)

// Definisi struct User
type User struct {
	Nama string
	Umur int
}

// Method dengan value receiver (tidak mengubah data asli)
func (u User) SayHello() {
	fmt.Printf("Halo, saya %s, umur saya %d tahun.\n", u.Nama, u.Umur)
}

// Method dengan pointer receiver (bisa ubah data asli)
func (u *User) UpdateProfile(namaBaru string, umurBaru int) {
	u.Nama = namaBaru
	u.Umur = umurBaru
}

// Method untuk menampilkan informasi user
func (u User) Tampilkan() {
	fmt.Println("==== Informasi User ====")
	fmt.Println("Nama:", u.Nama)
	fmt.Println("Umur:", u.Umur)
}

func main() {
	// Membuat user baru
	user1 := User{
		Nama: "Andi",
		Umur: 25,
	}

	// Panggil method SayHello
	user1.SayHello()

	// Tampilkan data awal
	user1.Tampilkan()

	// Update profile menggunakan method pointer
	fmt.Println("\nUpdate profil...")
	user1.UpdateProfile("Budi", 30)

	// Tampilkan data setelah update
	user1.Tampilkan()

	// Say hello lagi
	user1.SayHello()
}
