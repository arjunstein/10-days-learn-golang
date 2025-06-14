package main

import "fmt"

// struct user
type User struct {
	Name string
	Email string
	Age int
}

// method untuk user (value receiver)
func (u User) SayHello()  {
	fmt.Printf("Halo, nama saya %s. Email saya: %s\n", u.Name, u.Email)
}

// method untuk update umur (pointer receiver)
func (u *User) TambahUmur(tahun int)  {
	u.Age += tahun
}

// struct book
type Book struct {
	Title string
	Author string
	PageCount int
}

// method untuk menampilkan ringkasan
func (b Book) Summary()  {
	fmt.Printf("'%s' oleh %s (%d halaman)\n", b.Title, b.Author, b.PageCount)
}

// method untuk menambah halaman
func (b *Book) TambahHalaman(tambahan int)  {
	b.PageCount += tambahan
}

func main()  {
	fmt.Println("=== Struct & Method Demo ===")

	// === User ===
	user := User{
		Name: "Arjun",
		Email: "arjunstein@gmail.com",
		Age: 25,
	}

	user.SayHello()
	user.TambahUmur(3)
	fmt.Printf("Umur setelah ditambah: %d tahun\n", user.Age)

	// === Book ===
	buku := Book{
		Title: "Learn Golang",
		Author: "Arjun",
		PageCount: 200,
	}

	buku.Summary()
	buku.TambahHalaman(50)
	fmt.Println("Setelah ditambah halaman:")
	buku.Summary()
}