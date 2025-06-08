package Fungsi

import (
	"fmt"
	"strings"
)

func reamur(celcius float64) float64 {
	var reamur float64
	reamur = celcius * 4 / 5
	return reamur
}

func fahrenheit(celcius float64) float64 {
	var fahrenheit float64
	fahrenheit = celcius*9/5 + 32
	return fahrenheit
}

func kelvin(celcius float64) float64 {
	var kelvin float64
	kelvin = celcius + 273
	return kelvin
}

func Soal1fungsi() bool {
	fmt.Println(`
============================================
package main

import "fmt"

func reamur(celcius float64) float64 {
	var reamur float64
	reamur = celcius * 4 / 5
	return reamur
}

func fahrenheit(celcius float64) float64 {
	var fahrenheit float64
	fahrenheit = celcius*9/5 + 32
	return fahrenheit
}

func kelvin(celcius float64) float64 {
	var kelvin float64
	kelvin = celcius + 273
	return kelvin
}

func main() {
	var celcius_awal, celcius_akhir, step float64
	fmt.Scan(&celcius_awal, &celcius_akhir, &step)
	fmt.Printf("%-10s %-10s %-12s %-10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for celcius_awal <= celcius_akhir {
		fmt.Printf("%-10.2f %-10.2f %-12.2f %-10.2f\n", celcius_awal, reamur(celcius_awal), fahrenheit(celcius_awal), kelvin(celcius_awal))
		celcius_awal += step
	}
}
============================================`)
	var celcius_awal, celcius_akhir, step float64
	fmt.Scan(&celcius_awal, &celcius_akhir, &step)
	fmt.Printf("%-10s %-10s %-12s %-10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for celcius_awal <= celcius_akhir {
		fmt.Printf("%-10.2f %-10.2f %-12.2f %-10.2f\n", celcius_awal, reamur(celcius_awal), fahrenheit(celcius_awal), kelvin(celcius_awal))
		celcius_awal += step
	}
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
