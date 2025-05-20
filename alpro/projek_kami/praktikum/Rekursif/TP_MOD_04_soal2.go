package Rekursif

import "fmt"

func mod(a, b int) int {
	if b == 0 {
		return a
	} else {
		return mod(b, a%b)
	}

}

func soal2rekursif() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(mod(bil1, bil2))
}
