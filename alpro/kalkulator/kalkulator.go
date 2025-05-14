package kalkulator

import (
	"Project-Alpro/atribut"
	"fmt"
	"math"
)

var Choice2 string

const Narray = 10000

func MainMenu() {
	atribut.ClearScreen()
	kalkulatorSederhana()
}

func kalkulatorSederhana() {
	fmt.Println("Kalkulator Sederhana")
	var a, b int
	fmt.Println(`
		_____________________
		|  _________________  |
		| |              0. | |
		| |_________________| |
		|  ___ ___ ___   ___  |
		| | 7 | 8 | 9 | | + | |
		| |___|___|___| |___| |
		| | 4 | 5 | 6 | | - | |
		| |___|___|___| |___| |
		| | 1 | 2 | 3 | | x | |
		| |___|___|___| |___| |
		| | . | 0 | = | | ÷ | |
		| |___|___|___| |___| |
		|_____________________|
		
		1. Penjumlahan
		2. Pengurangan
		3. Perkalian
		4. Pembagian
		5. Pangkat
		6. Akar
		7. Modulus
		0. Keluar
   `)
	var pilihan string
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Penjumlahan: %d\n", penjumlahan(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "2":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Pengurangan: %d\n", pengurangan(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "3":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Perkalian: %d\n", perkalian(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "4":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Pembagian: %.2f\n", pembagian(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "5":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Pangkat: %d\n", pangkat(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "6":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Akar: %.2f\n", akar(a))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "7":
		fmt.Println("Masukkan angka pertama dan kedua: ")
		fmt.Scan(&a, &b)
		fmt.Printf("Hasil Modulus: %d\n", modulus(a, b))
		fmt.Println("Tetap di kalkulator atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "0":
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}
func penjumlahan(a, b int) int {
	return a + b
}
func pengurangan(a, b int) int {
	return a - b
}
func perkalian(a, b int) int {
	return a * b
}
func pembagian(a, b int) float64 {
	if b == 0 {
		fmt.Println("Pembagian dengan nol tidak terdefinisi")
		return 0
	}
	return float64(a) / float64(b)
}
func pangkat(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
func akar(a int) float64 {
	return math.Sqrt(float64(a))
}
func modulus(a, b int) int {
	return a % b
}
