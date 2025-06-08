package Tipebentukan

import (
	"fmt"
	"strings"
)

type mobil struct {
	merek                     string
	tahun_produksi, kecepatan int
}

func Soal1bentukan() bool {
	fmt.Println(`
============================================
package main

import "fmt"

type mobil struct {
	merek                     string
	tahun_produksi, kecepatan int
}

func main() {
	var m1, m2, m3 mobil
	var rata_rata_kecepatan float64
	fmt.Scan(&m1.merek, &m1.tahun_produksi, &m1.kecepatan)
	fmt.Scan(&m2.merek, &m2.tahun_produksi, &m2.kecepatan)
	fmt.Scan(&m3.merek, &m3.tahun_produksi, &m3.kecepatan)
	// hitung rata-rata kecepatan dari 3 mobil
	rata_rata_kecepatan = float64(m1.kecepatan+m2.kecepatan+m3.kecepatan) / 3
	// cetak data data mobil dengan rata-rata kecepatannya
	fmt.Printf("Rata-rata kecepatan mobil %s (%d), ", m1.merek, m1.tahun_produksi)
	fmt.Printf("mobil %s (%d), dan mobil %s (%d): ", m2.merek, m2.tahun_produksi, m3.merek, m3.tahun_produksi)
	fmt.Printf("%.2f\n", rata_rata_kecepatan)
}
============================================`)
	var m1, m2, m3 mobil
	var rata_rata_kecepatan float64
	fmt.Scan(&m1.merek, &m1.tahun_produksi, &m1.kecepatan)
	fmt.Scan(&m2.merek, &m2.tahun_produksi, &m2.kecepatan)
	fmt.Scan(&m3.merek, &m3.tahun_produksi, &m3.kecepatan)
	// hitung rata-rata kecepatan dari 3 mobil
	rata_rata_kecepatan = float64(m1.kecepatan+m2.kecepatan+m3.kecepatan) / 3
	// cetak data data mobil dengan rata-rata kecepatannya
	fmt.Printf("Rata-rata kecepatan mobil %s (%d), ", m1.merek, m1.tahun_produksi)
	fmt.Printf("mobil %s (%d), dan mobil %s (%d): ", m2.merek, m2.tahun_produksi, m3.merek, m3.tahun_produksi)
	fmt.Printf("%.2f\n", rata_rata_kecepatan)
	fmt.Println()
	var pilihan string
	// Variabel untuk mengontrol loop input validasi
	inputValid := false
	for !inputValid {
		fmt.Println("\n------------------------------------")
		fmt.Print("Ingin kembali ke menu utama (n)? ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		if strings.ToLower(pilihan) == "n" {
			inputValid = true // Input valid, keluar dari loop input
			// Mengembalikan true berarti kembali ke MainMenu
			return true // Kembali ke MainMenu
		} else {
			fmt.Println("Pilihan tidak valid. Harap masukkan 'n'.")
			// Loop akan terus berjalan sampai input 'n' diterima
		}
	}
	// Jika loop berakhir karena input valid, berarti 'n' sudah dimasukkan,
	return true
}
