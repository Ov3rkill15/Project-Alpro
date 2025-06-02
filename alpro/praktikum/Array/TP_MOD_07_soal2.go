package Array

import "fmt"

const NMAX1 int = 10

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}
func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
	}
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func Soal2array() {
	fmt.Println(`
============================================
package main

import "fmt"

const NMAX1 int = 10

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}
func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
	}
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func main() {
	var data tabInt
	var banyakData int
	bacaData(&data, &banyakData)
	fmt.Println()
	cetakData(data, banyakData)
}
============================================`)
	fmt.Println("SILAHKAN DICOBA: ")
	var data tabInt
	var banyakData int
	bacaData(&data, &banyakData)
	fmt.Println()
	cetakData(data, banyakData)
}
