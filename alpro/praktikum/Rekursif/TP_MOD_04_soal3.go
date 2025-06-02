package Rekursif

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return faktorial(n-1) * n
	}
}

func Soal3rekursif() {
	fmt.Println(`
============================================
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return faktorial(n-1) * n
	}
}

func main() {
	var bil int
	fmt.Scan(&bil)
	fmt.Println(faktorial(bil))

}
`)
	var bil int
	fmt.Scan(&bil)
	fmt.Println(faktorial(bil))

}
