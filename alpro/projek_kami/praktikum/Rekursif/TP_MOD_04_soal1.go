package Rekursif

import "fmt"

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func Soal1rekursif() {
	var bil int
	fmt.Scan(&bil)
	fmt.Println(fibo(bil))
}
