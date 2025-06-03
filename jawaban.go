// Buatlah sebuah fungsi Go bernama `hitungJumlah` yang menerima dua bilangan bulat ($a$ dan $b$) sebagai input, lalu mengembalikan hasil penjumlahan kedua bilangan tersebut. Cetak hasil pemanggilan fungsi ini dengan input 5 dan 3.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a,&b)
	hitungJumlah(a,b)
	fmt.Println(hitungJumlah(a,b))
}

func hitungJumlah(A int, B int)int{
	var nilai int
	nilai = A + b
	
}
