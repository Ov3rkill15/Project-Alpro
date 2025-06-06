package insertionsort

import (
	"fmt"
	"strings"
)

const nmax int = 99

type tabInt [nmax]int

func SoalInsertionSort() bool {
	fmt.Println(`
============================================
package insertionsort

import "fmt"

const nmax int = 99

type tabInt [nmax]int

func SoalInsertionSort() {
	var data tabInt
	var nData int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Println()
	cetakData(data, nData)
	insertionSort(&data, nData)
	cetakData(data, nData)
}

func insertionSort(A *tabInt, n int) {
	var i, pass int
	var temp int
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func bacaData(A *tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
============================================`)
	var data tabInt
	var nData int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Println()
	cetakData(data, nData)
	insertionSort(&data, nData)
	cetakData(data, nData)
	fmt.Println()
	var pilihan string
	// Variabel untuk mengontrol loop input validasi
	inputValid := false
	for !inputValid {
		fmt.Println("\n------------------------------------")
		fmt.Print("Ingin kembali ke menu utama (n)? ") // Asumsi setelah soal selesai, hanya ada opsi kembali
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
	// jadi kita mengembalikan true untuk kembali ke MainMenu.
	return true
}

func insertionSort(A *tabInt, n int) {
	var i, pass int
	var temp int
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func bacaData(A *tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
