package array

import (
	"Project-Alpro/atribut"
	"fmt"
)

var pilihan string
var pilihan2 string

func Submenu() {
	fmt.Println(`
==================
Selamat Datang di Pembelajaran Array
==================
1. Tata cara array,slice,map
2. soal array
3. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}

func MainMenu() {
	atribut.ClearScreen()
	Submenu()
	switch pilihan {
	case "1":
		atribut.ClearScreen()
		belajarArray()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}

	case "2":
		atribut.ClearScreen()
		soalArray()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "3":
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func belajarArray() {
	fmt.Print(`//deklarasi array
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy" // tidak bisa lebih dari 3 data jika di deklarasi hanya ada 3
	//atau dengan
	var num = [4]int{
		90, 80, 95, 90,
	}
	var kosong = [10]int{}

	//cara akses data dari array
	//1. panggil array
	fmt.Println(names) //[Eko Kurniawan Khannedy]
	fmt.Println(num)   // [90 80 95 90]
	fmt.Println()

	//2. pemanggilan data indeks 0 berupa data pertama
	fmt.Println(names[0]) //eko
	fmt.Println(num[0])   //90
	fmt.Println()

	//3. pemanggilan jumlah panjang array, bukan panjang data
	fmt.Println(len(names))  //3
	fmt.Println(len(num))    // 4
	fmt.Println(len(kosong)) //10bb
	fmt.Println()

	//4. mengubah value dari index tertentu
	names[0] = "Alwan"    // [Eko Kurniawan Khannedy] -> [Alwan Kurniawan Khannedy]
	num[1] = 100          // [90 100 95 90] -> [90 100 95 90]
	fmt.Println(names[0]) //Alwan
	fmt.Println(num[0])   // 100
	fmt.Println()

	fmt.Println("=======================SLICE===============================")
		KARAKTERISTIK ARRAY
		1.memiliki 3 tipe data, yaitu, pointer, length dan capacity
		2.pointer menunjuk data pertama di array
		3.length menunjuk panjang slice
		4.capacity adalah kapasitas dari slice, dimana length tidak boleh melebihi capacity
	
		var fruits = []string{"apple", "grape", "banana", "pisitan"}

		Ada beberapa cara membuat slice (partisi dari array)
		1. array[low:high]	(dari index low sampai sebelum high)
		2. array[low:] (dari index low sampai data akhir array)
		3. array [:high] (dari index 0 sampai sebelum high)
		4. array[:] () (dari index 0 sampai data akhir array)

	var fruits1 = fruits[1:3]
	var fruits2 = fruits[0:]
	var fruits3 = fruits[:3]
	var fruits4 = fruits[:]

	fmt.Println("array[low:high]")
	fmt.Println(cap(fruits1)) // 3
	fmt.Println(len(fruits1)) // 2
	fmt.Println()
	fmt.Println("array[low:]")
	fmt.Println(cap(fruits2)) // 4
	fmt.Println(len(fruits2)) // 4
	fmt.Println()
	fmt.Println("array[:high]")
	fmt.Println(cap(fruits3)) // 4
	fmt.Println(len(fruits3)) // 3
	fmt.Println()
	fmt.Println("array[:]")
	fmt.Println(cap(fruits4)) // 4
	fmt.Println(len(fruits4)) // 4

	//bukti bahwa slice adalah pointer
	fmt.Println(fruits)  // ["apple", "grape", "banana","pisitan"]
	fmt.Println(fruits1) // ["apple", "grape"]
	fruits[0] = "ALWAN"  // index [0] diganti dengan "ALWAN"
	fmt.Println(fruits)  // ["ALWAN", "grape","banana","pisitan"]
	fmt.Println(fruits1) // ["grape", "banana"]
	fmt.Println(fruits3) // ["ALWAN", "grape", "banana"]

		operasi slice
		1. len(array)
		2. cap(array)
		3. append(slice,data)
		4. make([]type data, length, capacity)
		5. copy(destination,source)

	// append(slice,data)
	fruits = append(fruits, "alwan") // ["apple", "grape", "banana", "pisitan", "alwan"]
	fmt.Println(fruits)              //jika capacity sudah penuh, maka akan membuat array baru

	// make([]type data, length, capacity)
	var data = make([]int, 5, 10)
	fmt.Println(data) // [0 0 0 0 0]

	// copy(destination,source)
	var data1 = []int{1, 2, 3, 4, 5}
	var data2 = []int{6, 7, 8, 9, 10}
	copy(data1, data2) // [6 7 8 9 10]
	fmt.Println(data1)
	data = append(data, append(data1, data2...)...)
	fmt.Println(data)
`)
}

func soalArray() {
	fmt.Println("COMING SOON!")
}
