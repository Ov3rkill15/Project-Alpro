package login

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

type Users struct {
	name     string
	Password string
}
type daftarNama [100]Users

func Login() {
	var username, password string
	var user daftarNama

	// Input username dan password
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println() // pindah baris setelah password

	if err != nil {
		fmt.Println("Gagal membaca password:", err)
		return
	}
	user[0] = Users{"admin", "rahasia"}
	user[1] = Users{"Alwan", "12345"}
	user[2] = Users{"Fathur", "12345"}

	for i := 0; i < len(user); i++ {
		if user[i].name == username && user[i].Password == string(passwordBytes) {
			fmt.Println("Login berhasil!")
			return
		}
	}
	password = string(passwordBytes)

	// Contoh validasi login
	if username == "admin" && password == "rahasia" {
		fmt.Println("Login berhasil!")
	} else {
		fmt.Println("Username atau password salah.")
	}
}
