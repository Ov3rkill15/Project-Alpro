package kalkulator

import "fmt"

func Menu() {
	fmt.Println("=== Kalkulator Sederhana ===")
	var a, b int
	fmt.Println("Masukkan angka pertama dan kedua: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Hasil penjumlahan: %d\n", a+b)
}
