package login

import (
	"Project-Alpro/atribut"
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/common-nighthawk/go-figure"
	"golang.org/x/term"
)

type user struct {
	Username, Password string
}

const nmax int = 45

type daftarNama [nmax]user

var users daftarNama
var jumlahPengguna int = 42

func Login(sign, helo *string) {
	var username, password string
	var signup user
	stop := true
	key := 0
	reader := bufio.NewReader(os.Stdin)
	welcome := figure.NewFigure("=== LOGIN ===", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")
	fmt.Printf("%15s\n", "1. Masuk")
	fmt.Printf("%15s\n", "2. Daftar")
	fmt.Print("pilih yang mana: ")
	signed, _ := reader.ReadString('\n')
	*sign = strings.TrimSpace(signed)

	if *sign == "Masuk" || *sign == "1" {
		welcome := figure.NewFigure("=== MASUK ===", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")
		for stop && key < 3 {
			fmt.Print("Masukkan Username: ")
			usernameInput, _ := reader.ReadString('\n')
			username = strings.TrimSpace(usernameInput)

			fmt.Print("Masukkan Password: ")
			passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca kata sandi:", err)
				return
			}
			password = string(passwordBytes)

			if authenticateUser(username, password) {
				fmt.Println("Login berhasil!")
				stop = false
				if helo != nil {
					*helo = username
				}
				atribut.ClearScreen()
				fmt.Println("Menuju Aplikasi")
				atribut.Loading(100)
				fmt.Println()
			} else {
				fmt.Println("Username atau password salah.")
				key++
				fmt.Printf("Tersisa %d kesempatan\n", 3-key)
			}
		}
	} else if *sign == "daftar" || *sign == "2" {
		welcome := figure.NewFigure("=== DAFTAR ===", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")
		if jumlahPengguna >= nmax {
			fmt.Println("Maaf, jumlah maksimum pengguna telah tercapai. Tidak bisa Daftar sekarang.")
		}

		fmt.Print("Masukkan Username: ")
		signup.Username, _ = reader.ReadString('\n')
		signup.Username = strings.TrimSpace(signup.Username)

		fmt.Print("Masukkan NIM: ")
		NIM, _ := reader.ReadString('\n')
		NIM = strings.TrimSpace(NIM)

		fmt.Println("Selamat datang,")
		fmt.Println("Nama:", signup.Username)
		fmt.Println("NIM:", NIM)

		for stop && key < 3 {
			fmt.Print("Masukkan Password: ")
			passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca kata sandi:", err)
			}

			fmt.Print("Konfirmasi Password: ")
			confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca konfirmasi password:", err)
			}

			if string(passwordBytes) != string(confirmBytes) {
				fmt.Println("Password dan konfirmasi tidak cocok.")
				key++
			} else {
				stop = false
				signup.Password = string(passwordBytes)
				users[jumlahPengguna] = signup
				jumlahPengguna++
				fmt.Println("Signup berhasil! Silakan login.")
			}
		}
	} else {
		fmt.Println("Pilihan tidak dikenali.")
	}
}

func authenticateUser(username, password string) bool {
	for i := 0; i < jumlahPengguna; i++ {
		if users[i].Username == username && users[i].Password == password {
			return true
		}
	}
	return false
}

func initUsers() {
	if nmax < 42 {
		fmt.Println("Peringatan: nmax terlalu kecil, hanya sebagian data akan dimuat.")
	}
	data := [42]user{
		{"admin", "admin123"},
		{"nathasyayuanmaharani", "0001"},
		{"theodoreelvisestrada", "0006"},
		{"dyahkusumawardani", "0009"},
		{"azrielraihaneldovahartoto", "0010"},
		{"muhammadilhamalifianda", "0022"},
		{"alyaazizaputeri", "0026"},
		{"ahmadabdansyakuro", "0029"},
		{"fathurrahmanalfarizi", "0035"},
		{"nuswantorosetyomukti", "0040"},
		{"anggitacahyatihidayat", "0041"},
		{"wibnuhijrahfranstio", "0048"},
		{"meyshaprimiandita", "0050"},
		{"muhamadfiqrihabibi", "0056"},
		{"fitriacahyani", "0060"},
		{"triansyahdaniswaraibrahim", "0062"},
		{"rakhaabdillahalkautsar", "0068"},
		{"avicenanaufallathif", "0073"},
		{"naylaassyifa", "0078"},
		{"williampetervanxnajoan", "0084"},
		{"rayvanalifarlomahesworo", "0087"},
		{"zaidansalamrojab", "0088"},
		{"audreyfredileyhanas", "0093"},
		{"muhammadnaelfadly", "0096"},
		{"nairacahayaputridarmawansinaga", "0100"},
		{"muhamadalwansuryadi", "0104"},
		{"dhafyahmadzubaidi", "0107"},
		{"muhammadfarisdhiyaylhaqsarbini", "0117"},
		{"nursyadira", "0123"},
		{"rayfitokrisnawijaya", "0124"},
		{"mochammadrafirisqullah", "0129"},
		{"iputugedeagastyakrisnawidartha", "0134"},
		{"rendil", "0137"},
		{"muhammadariqazzaki", "0138"},
		{"edmundyuliusgantur", "0155"},
		{"muhammadsayyidhuwaidi", "0157"},
		{"muhdzuljalalwaliikramjalil", "0160"},
		{"ramadhantangguhdefennder", "0003"},
		{"adzkiyaputrirahmawan", "0025"},
		{"fathimahradhiyya", "0029"},
		{"rakanghazianadiwjaya", "0034"},
		{"jihannabilamubarakah", "0037"},
	}

	// Hanya masukkan sesuai kapasitas array
	jumlahPengguna = 0
	for i := 0; i < nmax && i < len(data); i++ {
		users[i] = data[i]
		jumlahPengguna++
	}
}

func Mainlogin(sign, helo *string) {
	initUsers()
	for stop := true; stop; {
		Login(sign, helo)
		if *sign == "signin" {
			stop = false
		} else if *sign == "signup" {
			Login(sign, helo)
			*sign = "signup"
			stop = false
		}
	}
}
