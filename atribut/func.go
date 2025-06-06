package atribut

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Quiz struct {
	Nama                         string
	ID                           string
	TotalScore                   int
	GoLanguageScore              int
	PercabanganScore             int
	KonversiTipeDataScore        int
	OperasiMatematikaLogikaScore int
	PerulanganScore              int
	TipeDataGoScore              int
	VariableConstantScore        int
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Peringatan: Gagal membersihkan layar. Error: %v\n", err)
	}
}
func Loading(durationMs int) {
	fmt.Print("Loading")
	interval := 300 // interval antar titik (ms)
	dots := durationMs / interval

	for i := 0; i < dots; i++ {
		time.Sleep(time.Duration(interval) * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
}

// GetWaktuSekarang mengembalikan string representasi waktu saat ini (YYYY-MM-DD HH:MM:SS)
func GetWaktuSekarang() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func BukaPDF(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", path)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
	case "darwin": // MacOS
		cmd = exec.Command("open", path)
	default:
		return fmt.Errorf("sistem operasi tidak didukung")
	}

	return cmd.Start()
}
