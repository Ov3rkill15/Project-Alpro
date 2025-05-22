package kalkulator

import (
	"Project-Alpro/atribut"
	"fmt"
	"math"
)

var Choice2 string
var previousValue *int

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
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan angka kedua yang ingin dijumlahkan dengan nilai sebelumnya: ")
			fmt.Scan(&b)
			result := penjumlahan(*previousValue, b)
			previousValue = &result
			fmt.Printf("Hasil Penjumlahan: %d\n", result)
		} else {
			fmt.Print("Masukkan angka pertama dan kedua: ")
			fmt.Scan(&a, &b)
			result := penjumlahan(a, b)
			previousValue = &result
			fmt.Printf("Hasil Penjumlahan: %d\n", result)
		}
		tanyaKembali()
	case "2":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan angka kedua yang ingin dikurangkan dari nilai sebelumnya: ")
			fmt.Scan(&b)
			result := pengurangan(*previousValue, b)
			previousValue = &result
			fmt.Printf("Hasil Pengurangan: %d\n", result)
		} else {
			fmt.Print("Masukkan angka pertama dan kedua: ")
			fmt.Scan(&a, &b)
			result := pengurangan(a, b)
			previousValue = &result
			fmt.Printf("Hasil Pengurangan: %d\n", result)
		}
		tanyaKembali()
	case "3":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan angka kedua yang ingin dikalikan dengan nilai sebelumnya: ")
			fmt.Scan(&b)
			result := perkalian(*previousValue, b)
			previousValue = &result
			fmt.Printf("Hasil Perkalian: %d\n", result)
		} else {
			fmt.Print("Masukkan angka pertama dan kedua: ")
			fmt.Scan(&a, &b)
			result := perkalian(a, b)
			previousValue = &result
			fmt.Printf("Hasil Perkalian: %d\n", result)
		}
		tanyaKembali()
	case "4":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan angka kedua yang ingin digunakan sebagai pembagi nilai sebelumnya: ")
			fmt.Scan(&b)
			resultFloat := pembagian(*previousValue, b)
			if b != 0 {
				result := int(resultFloat)
				previousValue = &result
			}
			fmt.Printf("Hasil Pembagian: %.2f\n", resultFloat)
		} else {
			fmt.Print("Masukkan angka pertama dan kedua: ")
			fmt.Scan(&a, &b)
			resultFloat := pembagian(a, b)
			if b != 0 {
				result := int(resultFloat)
				previousValue = &result
			}
			fmt.Printf("Hasil Pembagian: %.2f\n", resultFloat)
		}
		tanyaKembali()
	case "5":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan pangkat yang ingin diterapkan pada nilai sebelumnya: ")
			fmt.Scan(&b)
			result := pangkat(*previousValue, b)
			previousValue = &result
			fmt.Printf("Hasil Pangkat: %d\n", result)
		} else {
			fmt.Print("Masukkan angka dan pangkat: ")
			fmt.Scan(&a, &b)
			result := pangkat(a, b)
			previousValue = &result
			fmt.Printf("Hasil Pangkat: %d\n", result)
		}
		tanyaKembali()
	case "6":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			result := akar(*previousValue)
			prevInt := int(result)
			previousValue = &prevInt
			fmt.Printf("Hasil Akar: %.2f\n", result)
		} else {
			fmt.Print("Masukkan angka: ")
			fmt.Scan(&a)
			result := akar(a)
			prevInt := int(result)
			previousValue = &prevInt
			fmt.Printf("Hasil Akar: %.2f\n", result)
		}
		tanyaKembali()
	case "7":
		if previousValue != nil {
			fmt.Printf("Nilai sebelumnya: %d\n", *previousValue)
			fmt.Print("Masukkan angka kedua untuk operasi modulus dengan nilai sebelumnya: ")
			fmt.Scan(&b)
			result := modulus(*previousValue, b)
			previousValue = &result
			fmt.Printf("Hasil Modulus: %d\n", result)
		} else {
			fmt.Print("Masukkan angka pertama dan kedua: ")
			fmt.Scan(&a, &b)
			result := modulus(a, b)
			previousValue = &result
			fmt.Printf("Hasil Modulus: %d\n", result)
		}
		tanyaKembali()
	case "0":
		fmt.Println("Terima kasih telah menggunakan kalkulator!")
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		kalkulatorSederhana()
	}
}

func tanyaKembali() {
	fmt.Print("Tetap di kalkulator atau kembali ke menu utama? (y/n): ")
	fmt.Scan(&Choice2)
	if Choice2 == "y" || Choice2 == "Y" {
		MainMenu()
	} else {
		fmt.Println("Terima kasih telah menggunakan kalkulator!")
		return
	}
}
func penjumlahan(a, b int) int {
	var hasil int
	var angka int

	hasil = a + b
	fmt.Printf("Hasil sementara: %d\n", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = hasil + angka
		fmt.Printf("Hasil sementara: %d\n", hasil)
	}

	return hasil
}
func pengurangan(a, b int) int {
	var hasil int
	var angka int

	hasil = a - b
	fmt.Printf("Hasil sementara: %d\n", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = hasil - angka
		fmt.Printf("Hasil sementara: %d\n", hasil)
	}

	return hasil
}

func perkalian(a, b int) int {
	var hasil int
	var angka int

	hasil = a * b
	fmt.Printf("Hasil sementara: %d\n", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = hasil * angka
		fmt.Printf("Hasil sementara: %d\n", hasil)
	}

	return hasil
}

func pembagian(a, b int) float64 {
	var hasil float64
	var angka float64
	if b == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak terdefinisi")
		return 0
	}
	hasil = float64(a) / float64(b)
	fmt.Println("Hasil sementara : ", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = hasil / angka
		fmt.Println("Hasil sementara:", hasil)
	}

	return (hasil)
}

func pangkat(a, b int) int {
	var angka int

	if b < 0 {
		fmt.Println("Error: Pangkat negatif tidak didukung untuk integer")
		return 0
	}

	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}

	fmt.Printf("Hasil awal: %d^%d = %d\n", a, b, result)

	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasilLama := result
		result = 1
		for i := 0; i < angka; i++ {
			result *= hasilLama
		}

		fmt.Printf("Hasil: %d^%d = %d\n", hasilLama, angka, result)
	}

	return result
}

func akar(a int) float64 {
	var hasil float64
	var angka float64
	if a < 0 {
		fmt.Println("Error: Akar dari bilangan negatif tidak terdefinisi dalam bilangan real")
		return 0
	}
	hasil = math.Sqrt(float64(a))
	fmt.Println("Hasil sementara : ", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = math.Sqrt(float64(hasil))
		fmt.Println("Hasil sementara:", hasil)
	}

	return (hasil)

}

func modulus(a, b int) int {
	var hasil int
	var angka int
	if b == 0 {
		fmt.Println("Error: Modulus dengan nol tidak terdefinisi")
		return 0
	}
	hasil = a % b
	fmt.Println("Hasil sementara : ", hasil)
	for {
		fmt.Print("Masukkan angka berikutnya (0 untuk selesai): ")
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}
		hasil = hasil % angka
		fmt.Println("Hasil sementara:", hasil)
	}
	return hasil

}
