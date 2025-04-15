package kalkulator

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

func MainMenu() {
	clearScreen()
	Submenu()
	fmt.Print("Masukkan pilihan: ")
	var pilihan string
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1":
		clearScreen()
		kalkulatorSederhana()
	case "2":
		fmt.Println("GUI belum tersedia.")
	case "3":
		fmt.Println("Terima kasih telah menggunakan kalkulator ini.")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Submenu() {
	fmt.Println(`
=================
Mau Kalkulator mana?
=================
1. Sederhana
2. GUI
3. Keluar
	`)
}

func kalkulatorSederhana() {
	fmt.Println("Kalkulator Sederhana")
	var a, b int
	fmt.Println("Masukkan angka pertama dan kedua: ")
	fmt.Scan(&a, &b)
	fmt.Println("Silahkan pilih operasi:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Pangkat (a^b)")
	fmt.Println("6. Akar")
	fmt.Println("7. Modulus")
	fmt.Println("8. Faktorial")
	fmt.Println("9. Kombinasi")
	fmt.Println("10. Permutasi")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ")
	var pilihan string
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1":
		fmt.Printf("Hasil Penjumlahan: %d\n", penjumlahan(a, b))
	case "2":
		fmt.Printf("Hasil Pengurangan: %d\n", pengurangan(a, b))
	case "3":
		fmt.Printf("Hasil Perkalian: %d\n", perkalian(a, b))
	case "4":
		fmt.Printf("Hasil Pembagian: %.2f\n", pembagian(a, b))
	case "5":
		fmt.Printf("Hasil Pangkat: %d\n", pangkat(a, b))
	case "6":
		fmt.Printf("Hasil Akar: %.2f\n", akar(a))
	case "7":
		fmt.Printf("Hasil Modulus: %d\n", modulus(a, b))
	case "0":
		fmt.Println("Terima kasih telah menggunakan kalkulator ini.")
		os.Exit(0)
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
