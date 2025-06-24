package main

import "fmt"

func main()  {
	// slice of map
	mahasiswa := []map[string]string{
		{"nama": "Budi", "jurusan": "Teknik Informatika"},
		{"nama": "Siti", "jurusan": "Sistem Informasi"},
		{"nama": "Andi", "jurusan": "Teknik Komputer"},
	}

	// add new data
	mahasiswa = append(mahasiswa, map[string]string{"nama": "Dewi", "jurusan": "Teknik Elektro"})

	// print all data
	fmt.Println("Data Mahasiswa:")
	for i, mhs := range mahasiswa {
		fmt.Printf("Mahasiswa %d: Nama: %s, Jurusan: %s\n", i+1, mhs["nama"], mhs["jurusan"])
	}
}