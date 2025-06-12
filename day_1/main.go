package main

import "fmt"

const AppName = "Golang Day 1"

func main(){
	// 1. hello world
	fmt.Println("Selamat data di", AppName)

	// 2. variabel
	var nama string = "Arjun"
	umur := 25
	var tinggi float64 = 162.5
	aktif := true
	
	fmt.Println("Nama", nama)
	fmt.Println("Umur", umur)
	fmt.Println("Tinggi", tinggi)
	fmt.Println("Aktif", aktif)

	// 3. array dan slice
	angka := [3]int{1,2,3} // array
	hobi := []string{"coding","ngopi","traveling"} // slice

	fmt.Println("Angka", angka)
	fmt.Println("Hobi", hobi)

	// 4. Map
	user := map[string]string{
		"username":"Arjunaja",
		"email":"arjun@gmail.com",
	}

	fmt.Println("User", user["username"], "-", user["email"])

	// 5. input dari pengguna
	var a, b int
	fmt.Print("Masukan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukan angka kedua: ")
	fmt.Scan(&b)

	hasil := a + b
	fmt.Printf("Hasil penjumlahan %d + %d = %d\n", a, b, hasil)
}