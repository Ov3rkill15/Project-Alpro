package Procedure

import "fmt"

func hitungMenang(g, k int, jumlah_menang *int) {
	if g > k {
		*jumlah_menang++
	}
}
func hitungDraw(g, k int, jumlah_draw *int) {
	if g == k {
		*jumlah_draw++
	}
}
func hitungKalah(g, k int, jumlah_kalah *int) {
	if g < k {
		*jumlah_kalah++
	}
}
func Goal(g, k int, jumlah_goal, jumlah_bobol, selisih_goal *int) {
	*jumlah_goal += g
	*jumlah_bobol += k
	*selisih_goal = *jumlah_goal - *jumlah_bobol
}
func hitungPoint(jp *int) {
	*jp = jm*3 + jd
}

var jm, jd, jk, jg, jb, sg, poin int // dideklarasi secara global

func Soal2prosedur() {
	fmt.Println(`
============================================
package main

import "fmt"

func hitungMenang(g, k int, jumlah_menang *int) {
	if g > k {
		*jumlah_menang++
	}
}
func hitungDraw(g, k int, jumlah_draw *int) {
	if g == k {
		*jumlah_draw++
	}
}
func hitungKalah(g, k int, jumlah_kalah *int) {
	if g < k {
		*jumlah_kalah++
	}
}
func Goal(g, k int, jumlah_goal, jumlah_bobol, selisih_goal *int) {
	*jumlah_goal += g
	*jumlah_bobol += k
	*selisih_goal = *jumlah_goal - *jumlah_bobol
}
func hitungPoint(jp *int) {
	*jp = jm*3 + jd
}

var jm, jd, jk, jg, jb, sg, poin int // dideklarasi secara global

func main() {
	var match, goal, bobol int

	fmt.Scan(&match)
	for i := 0; i < match; i++ {
		fmt.Scan(&goal, &bobol)
		hitungMenang(goal, bobol, &jm)
		hitungDraw(goal, bobol, &jd)
		hitungKalah(goal, bobol, &jk)
		Goal(goal, bobol, &jg, &jb, &sg)
		hitungPoint(&poin)

	}
	fmt.Println(match, jm, jd, jk, jg, jb, sg, poin)
}
============================================`)
	var match, goal, bobol int

	fmt.Scan(&match)
	for i := 0; i < match; i++ {
		fmt.Scan(&goal, &bobol)
		hitungMenang(goal, bobol, &jm)
		hitungDraw(goal, bobol, &jd)
		hitungKalah(goal, bobol, &jk)
		Goal(goal, bobol, &jg, &jb, &sg)
		hitungPoint(&poin)

	}
	fmt.Println(match, jm, jd, jk, jg, jb, sg, poin)
}
