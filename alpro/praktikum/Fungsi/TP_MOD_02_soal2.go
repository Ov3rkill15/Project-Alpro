package Fungsi

import (
	"fmt"
	"unicode"
)

func lowToUpper(char rune) rune {
	return unicode.ToUpper(char)
}
func Soal2fungsi() {
	fmt.Println(`
============================================
package main

import (
	"fmt"
	"unicode"
)

func lowToUpper(char rune) rune {
	return unicode.ToUpper(char)
}
func main() {
	var input rune
	fmt.Scanf("%c", &input)
	fmt.Printf("%c\n", lowToUpper(input))

}
============================================`)
	var input rune
	fmt.Scanf("%c", &input)
	fmt.Printf("%c\n", lowToUpper(input))

}
