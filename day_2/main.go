package main

import (
	"errors"
	"fmt"
)

// fungsi dasar operasi
func tambah(a, b int) int  {
	return a + b
}

func kurang(a, b int) int  {
	return a - b
}

func kali(a, b int) int  {
	return a * b
}

func bagi(a, b int) (int, error)  {
	if b == 0 {
		return 0, errors.New("error: tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// fungsi penilaian dengan if/else
func nilaiGrade(nilai int) string  {
	if nilai >= 90 {
		return "A"
	} else if nilai >= 80 {
		return "B"
	} else if nilai >= 70 {
		return "C"
	} else {
		return "D atau E"
	}
}

// fungsi penilaian hari dengan switch
func infoHari(hari string)  {
	switch hari {
	case "senin":
		fmt.Println("Awal minggu, semangat kerja!")
	case "sabtu", "minggu":
		fmt.Println("Akhir pekan, waktunya traveling")
	default:
		fmt.Println("Hari biasa")
	}
}

// fungsi looping dan menampilkan daftar
func tampilkanHobi(hobi []string)  {
	fmt.Println("Daftar hobi: ")
	for i, h := range hobi {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

func main()  {
	// if, switch dan loop
	fmt.Println("=== Demo If/Switch/Loop ===")
	grade := nilaiGrade(87)
	fmt.Println("Nilai 87 mendapatkan grade: ", grade)

	infoHari("minggu")

	hobi := []string{"ngoding", "browsing", "bisnis"}
	tampilkanHobi(hobi)

	// kalkulator CLI
	fmt.Println("\n=== Kalkulator CLI ===")
	var a, b, pilihan int

	fmt.Print("Masukan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukan angka kedua: ")
	fmt.Scan(&b)

	fmt.Println("1. Tambah (+)")
	fmt.Println("2. Kurang (-)")
	fmt.Println("3. Kali (*)")
	fmt.Println("4. Bagi (/)")
	fmt.Print("Pilih operasi [1-4]: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println("Hasil: ", tambah(a, b))
	case 2:
		fmt.Println("Hasil: ", kurang(a, b))
	case 3:
		fmt.Println("Hasil: ", kali(a, b))
	case 4:
		hasil, err := bagi(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Hasil: ", hasil)
		}
	default:
		fmt.Println("Pilihan tidak valid")
	}
}