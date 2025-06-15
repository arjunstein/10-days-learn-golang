package main

import "fmt"

func main()  {
	// soal 1.1
	var nama string = "Arjun"
	var umur int = 25
	var tinggi float64 = 162.5
	var isBelajarGolang bool = true

	fmt.Println("== soal 1.1 ===")
	fmt.Println("Nama: ", nama)
	fmt.Println("Umur: ", umur)
	fmt.Println("Tinggi: ", tinggi)
	fmt.Println("Sedang belajar Golang? ", isBelajarGolang)

	// soal 1.2
	var x = 16
	var y = 5

	fmt.Println("== soal 1.2 ===")
	fmt.Println(x / y)
	fmt.Println(float64(x) / float64(y))

	// soal 2.1
	const phi = 3.14
	var radius = 7

	fmt.Println("== soal 2.1 ===")
	luas := phi * float64(radius) * float64(radius)
	fmt.Println("Luas lingkaran: ", luas)

	// soal 3.1
	var guy string
	var age int

	fmt.Println("=== soal 3.1 ===")
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&guy)
	fmt.Print("Masukan umur anda: ")
	fmt.Scan(&age)
	fmt.Printf("Halo, %s! usia kamu %d tahun.\n", guy, age)

	// soal 4.1
	a := 7
	b := 3

	fmt.Println("=== soal 4.1 ===")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// soal 5.1
	var harga int = 15000
	var diskon float64 = 0

	if harga > 10000 {
		diskon = 0.10 // 10%
	}

	potongan := float64(harga) * diskon
	total := float64(harga) - potongan

	fmt.Println("Harga awal: ", harga)
	fmt.Printf("Diskon: %.0f%%\n", diskon * 100)
	fmt.Printf("Total bayar: Rp. %.0f\n", total)
}