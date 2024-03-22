package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk merepresentasikan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Map untuk menyimpan data teman berdasarkan nomor absen
var dataTeman = map[int]Teman{
	1: {"John", "Jl. Kenanga No. 123", "Pengembang Perangkat Lunak", "Ingin belajar lebih banyak tentang Go."},
	2: {"Alice", "Jl. Mawar No. 456", "Desainer Grafis", "Ingin meningkatkan keterampilan dalam pemrograman."},
	3: {"Bob", "Jl. Anggrek No. 789", "Data Analyst", "Ingin mengembangkan aplikasi web."},
	// Tambahkan data teman lain di sini sesuai kebutuhan
}

// Fungsi untuk menampilkan data teman berdasarkan nomor absen
func showData(absen int) {
	teman, found := dataTeman[absen]
	if found {
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golajng:", teman.Alasan)
	} else {
		fmt.Println("Data teman dengan absen", absen, "tidak ditemukan.")
	}
}

func main() {
	// Mendapatkan argumen nomor absen dari command line
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	// Konversi argumen ke dalam bentuk integer
	absenStr := os.Args[1]
	absen, err := strconv.Atoi(absenStr)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	// Menampilkan data teman berdasarkan nomor absen yang diberikan
	showData(absen)
}
