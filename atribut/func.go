package atribut

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()
}

func Loading(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Print(".")
	}
}

func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin": // macOS
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		fmt.Println("Gagal membuka browser:", err)
	}
}
