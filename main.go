package main

import (
	"Project-Alpro/admin"
	"Project-Alpro/alpro"
	"Project-Alpro/login"

	// "Project-Alpro/linklanjutan"
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	var input string
	var stop bool = false // Variabel untuk menghentikan loop
	var sign, helo string
	login.Mainlogin(&sign, &helo)
	if helo == "admin" && sign == "1" || sign == "masuk" {
		admin.MainAdmin()
	}
	switch helo {
	case "nathasyayuanmaharani":
		helo = "Yuan"
	case "theodoreelvisestrada":
		helo = "Elvis"
	case "dyahkusumawardani":
		helo = "Dyah"
	case "azrielraihaneldovahartoto":
		helo = "Azriel Jids"
	case "muhammadilhamalifianda":
		helo = "Ilham"
	case "alyaazizaputeri":
		helo = "Alya"
	case "ahmadabdansyakuro":
		helo = "Abdan"
	case "fathurrahmanalfarizi":
		helo = "Fathur"
	case "nuswantorosetyomukti":
		helo = "Nuswantoro"
	case "anggitacahyatihidayat":
		helo = "Anggita"
	case "wibnuhijrahfranstio":
		helo = "Wibnu"
	case "meyshaprimiandita":
		helo = "Meysha"
	case "muhamadfiqrihabibi":
		helo = "Fiqri"
	case "fitriacahyani":
		helo = "Fitria"
	case "triansyahdaniswaraibrahim":
		helo = "Triansyah"
	case "rakhaabdillahalkautsar":
		helo = "Rakha"
	case "avicenanaufallathif":
		helo = "Avicenna"
	case "naylaassyifa":
		helo = "Nayla"
	case "williampetervanxnajoan":
		helo = "William"
	case "rayvanalifarlomahesworo":
		helo = "Rayvan"
	case "zaidansalamrojab":
		helo = "Zaidan"
	case "audreyfredileyhanas":
		helo = "Audrey"
	case "muhammadnaelfadly":
		helo = "Nael"
	case "nairacahayaputridarmawansinaga":
		helo = "Naira"
	case "muhamadalwansuryadi":
		helo = "Alwan"
	case "dhafyahmadzubaidi":
		helo = "Dhafy"
	case "muhammadfarisdhiyaylhaqsarbini":
		helo = "Faris"
	case "nursyadira":
		helo = "Nora"
	case "rayfitokrisnawijaya":
		helo = "Fito"
	case "mochammadrafirisqullah":
		helo = "Rafi"
	case "iputugedeagastyakrisnawidartha":
		helo = "Agastya"
	case "rendil":
		helo = "Rendil"
	case "muhammadariqazzaki":
		helo = "Ariq"
	case "edmundyuliusgantur":
		helo = "Edmund"
	case "muhammadsayyidhuwaidi":
		helo = "Sayyit kumis"
	case "muhdzuljalalwaliikramjalil":
		helo = "Dzul"
	case "ramadhantangguhdefennder":
		helo = "SENSEI"
	case "adzkiyaputrirahmawan":
		helo = "Adzkiya"
	case "fathimahradhiyya":
		helo = "Radhi"
	case "rakanghazianadiwjaya":
		helo = "Rakan"
	case "jihannabilamubarakah":
		helo = "Jihan"
	}

	if sign == "masuk" || sign == "1" {
		menuKonfirmasi("ya")
	} else {
		menuKonfirmasi("tidak")
	}

	for !stop {
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan menu
		user := figure.NewFigure(helo, "doom", true).String()
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		fmt.Print("\033[32m")
		fmt.Print(user)
		fmt.Print("\033[0m")
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			atribut.ClearScreen()
			fmt.Println("Menuju Pengenalan Pemrograman...")
			atribut.Loading(100)
			pengpro.MainMenu()
		case "2":
			atribut.ClearScreen()
			fmt.Println("Menuju Algoritma Pemrograman...")
			atribut.Loading(100)
			alpro.MainMenu()
		case "3":
			atribut.ClearScreen()
			login.Mainlogin(&sign, &helo)
		case "10":
			atribut.ClearScreen()
			fmt.Println("COMING SOON!!!")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading(100)
		case "0":
			stop = true
			atribut.ClearScreen()
			figure.NewFigure("THX THX THX", "basic", true).Blink(5000, 1000, 500)

		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(100)
			fmt.Println()
		}

	}

}

func menuKonfirmasi(n string) {
	// var konfirmasi string
	var konfirmasi2 string
	switch n {
	case "ya":
		atribut.ClearScreen()
		fmt.Println("Kami sarankan untuk kamu langsung menuju pilihan 2")
		fmt.Println("1. Setuju")
		fmt.Println("2. Dari awal aja!")
		fmt.Scanln(&konfirmasi2)
		cek := strings.ToLower(konfirmasi2)
		switch cek {
		case "1", "setuju":
			fmt.Println("Menuju Algoritma Pemrograman!")
			atribut.Loading(100)
			atribut.ClearScreen()
			alpro.MainMenu()
		case "2", "dari awal aja", "awal", "dari awal":
			fmt.Println("Oke, mari belajar sama-sama!")
			atribut.Loading(100)
			atribut.ClearScreen()
			menu()
		}
	case "tidak":
		fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Apakah kamu pernah belajar pemrograman go?
=========================================
1. ya
2. tidak
=========================================`)
		fmt.Scan(&konfirmasi2)
		cek := strings.ToLower(konfirmasi2)
		switch cek {
		case "1", "setuju":
			fmt.Println("Menuju Algoritma Pemrograman!")
			atribut.Loading(100)
			atribut.ClearScreen()
			alpro.MainMenu()
		case "2", "dari awal aja", "awal", "dari awal":
			fmt.Println("Oke, mari belajar sama-sama!")
			atribut.Loading(100)
			atribut.ClearScreen()
			menu()
		}
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}

}

func menu() {
	letsgo := figure.NewFigure("LET'S GO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(letsgo)     // Cetak teks ASCII
	fmt.Print("\033[0m")  // Reset warna ke default
	fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Materi yang tersedia:
=========================================
1. Pengenalan Pemrograman
2. Algoritma dan Pemrograman
3. Logout


=========================================
10. about us:
=========================================

0. keluar
	`)
}
