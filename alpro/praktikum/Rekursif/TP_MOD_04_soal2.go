package Rekursif

import "fmt"

func mod(a, b int) int {
	if b == 0 {
		return a
	} else {
		return mod(b, a%b)
	}

}

func Soal2rekursif() {
	fmt.Println(`
============================================
package main

import "fmt"

func mod(a, b int) int {
	if b == 0 {
		return a
	} else {
		return mod(b, a%b)
	}

}

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(mod(bil1, bil2))
}

============================================
		`)
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(mod(bil1, bil2))
}
