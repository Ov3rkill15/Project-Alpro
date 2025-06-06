package SelectionSort

import (
	"fmt"
	"strings"
)

const NMAX int = 100

type tabInt [NMAX]int

func SoalSelectionSort() bool {
	fmt.Println(`
============================================
package main

import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Println()
	cetakData(data, nData)
	SelectionSort(&data, nData)
	cetakData(data, nData)

}

func SelectionSort(A *tabInt, N int) {
	var i, idx, pass int
	var temp int
	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i] > A[idx] {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
func bacaData(A *tabInt, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

============================================
		`)
	var data tabInt
	var nData int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Println()
	cetakData(data, nData)
	SelectionSort(&data, nData)
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

func SelectionSort(A *tabInt, N int) {
	var i, idx, pass int
	var temp int
	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i] > A[idx] {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
func bacaData(A *tabInt, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
