package kalkulator

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

func MainMenu() {
	clearScreen()
	fmt.Println("Kalkulator Sederhana")
	fmt.Println("1. Kalkulator Sederhana")
	fmt.Println("2. GUI")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan: ")
	var pilihan string
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1":
		kalkulatorSederhana()
	case "2":
		fmt.Println("GUI belum tersedia.")
	case "3":
		fmt.Println("Terima kasih telah menggunakan kalkulator ini.")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
	fmt.Println("Tekan Enter untuk kembali ke menu...")
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
	case "8":
		fmt.Printf("Hasil Faktorial: %d\n", faktorial(a))
	case "9":
		fmt.Printf("Hasil Kombinasi: %d\n", kombinasi(a, b))
	case "10":
		fmt.Printf("Hasil Permutasi: %d\n", permutasi(a, b))
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
func faktorial(a int) int {
	if a == 0 || a == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	return result
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
