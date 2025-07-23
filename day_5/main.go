package main

import (
	"errors"
	"fmt"
)

// Fungsi dasar
func tambah(a, b int) int {
	return a + b
}

func kurang(a, b int) int {
	return a - b
}

func kali(a, b int) int {
	return a * b
}

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// Variadic function: total dari semua nilai
func total(nilai ...int) int {
	jumlah := 0
	for _, v := range nilai {
		jumlah += v
	}
	return jumlah
}

// Variadic function: rata-rata, dengan error jika tidak ada data
func rataRata(nilai ...int) (float64, error) {
	if len(nilai) == 0 {
		return 0, errors.New("data kosong, tidak bisa menghitung rata-rata")
	}
	total := 0
	for _, v := range nilai {
		total += v
	}
	return float64(total) / float64(len(nilai)), nil
}

func main() {
	// Demo fungsi biasa
	fmt.Println("=== Operasi Dasar ===")
	a, b := 10, 5
	fmt.Println("Tambah:", tambah(a, b))
	fmt.Println("Kurang:", kurang(a, b))
	fmt.Println("Kali  :", kali(a, b))

	hasilBagi, err := bagi(a, 0)
	if err != nil {
		fmt.Println("Bagi  :", err)
	} else {
		fmt.Println("Bagi  :", hasilBagi)
	}

	// Demo variadic function
	fmt.Println("\n=== Total dan Rata-rata ===")
	nilai := []int{80, 75, 95}
	fmt.Println("Nilai:", nilai)
	fmt.Println("Total nilai:", total(nilai...))

	rata, err := rataRata(nilai...)
	if err != nil {
		fmt.Println("Rata-rata:", err)
	} else {
		fmt.Printf("Rata-rata: %.2f\n", rata)
	}

	// Coba dengan data kosong
	fmt.Println("\n=== Coba Rata-rata Kosong ===")
	_, errKosong := rataRata()
	if errKosong != nil {
		fmt.Println("Error:", errKosong)
	}
}
