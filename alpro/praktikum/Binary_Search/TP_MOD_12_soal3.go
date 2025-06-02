package Binsearch

import "fmt"

const N = 10

type tabInt [N]int

func Soal3BinSearch() {
	var data tabInt
	var nData, x1 int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Scan(&x1)
	fmt.Println("Data dalam array:")
	cetakData(data, nData)
	if binSearch3(data, nData, x1) == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data ditemukan pada index ke-", binSearch3(data, nData, x1))
	}
}

func bacaData(data *tabInt, nData int) {
	for i := 0; i < nData; i++ {
		fmt.Scan(&data[i])
	}
}
func cetakData(data tabInt, nData int) {
	for i := 0; i < nData; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func binSearch3(tab tabInt, n, x int) int {
	fmt.Println(`
============================================
package main

import "fmt"

const N = 10

type tabInt [N]int

func main() {
	var data tabInt
	var nData, x1 int
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Scan(&x1)
	fmt.Println("Data dalam array:")
	cetakData(data, nData)
	if binSearch3(data, nData, x1) == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data ditemukan pada index ke-", binSearch3(data, nData, x1))
	}
}

func bacaData(data *tabInt, nData int) {
	for i := 0; i < nData; i++ {
		fmt.Scan(&data[i])
	}
}
func cetakData(data tabInt, nData int) {
	for i := 0; i < nData; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func binSearch3(tab tabInt, n, x int) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && tab[mid] != x {
		mid = (left + right) / 2
		if x < tab[mid] {
			right = mid - 1
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}
============================================`)
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && tab[mid] != x {
		mid = (left + right) / 2
		if x < tab[mid] {
			right = mid - 1
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}
