package Procedure

import (
	"fmt"
)

func menu() {
	fmt.Println("------------------")
	fmt.Println("Menu Kalkulator:")
	fmt.Println("------------------")
	fmt.Println("1. Hitung Jumlah")
	fmt.Println("2. Hitung Kali")
	fmt.Println("3. Hitung Bagi")
	fmt.Println("4. Keluar")
	fmt.Println("------------------")
}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukkan dua angka untuk dijumlahkan: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("Hasil: %.d\n", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukkan dua angka untuk dikalikan: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("Hasil: %.d\n", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukkan dua angka untuk dibagi (angka1 angka2): ")
	fmt.Scanln(&a, &b)
	if b == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
	} else {
		fmt.Printf("Hasil: %.1f\n", a/b)
	}
}

func soal3prosedur() {
	var pilih int

	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
