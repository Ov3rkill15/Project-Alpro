package praktikum

import (
	"fmt"
)

var pilihan string

func MainMenu() {
	fmt.Println(`
Berikut Adalah Pilihan Tugas Pendahuluan:
1. Fungsi
2. Prosedur
3. Rekursif
4. Tipe Bentukan
5. Array
6. Max/Min Searching
7. Sequential Search
8. Binary Search
9. Selection Sort
	`)
	fmt.Scan(&pilihan)

	switch pilihan {
	case "1":
		fmt.Println("Fungsi")
	case "2":
		fmt.Println("Prosedur")
	case "3":
		fmt.Println("Rekursif")
	case "4":
		fmt.Println("Tipe Bentukan")
	case "5":
		fmt.Println("Array")
	case "6":
		fmt.Println("Max/Min Searching")
	case "7":
		fmt.Println("Sequential Search")
	case "8":
		fmt.Println("Binary Search")
	case "9":
		fmt.Println("Selection Sort")

	}
}
