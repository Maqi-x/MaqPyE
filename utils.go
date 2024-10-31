package main

import (
	"runtime"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"os/exec"
	"log"
	"path/filepath")

func runPythonScript(path string) {
	dir := filepath.Dir(path)
    cmd := exec.Command("x-terminal-emulator", "-e", "bash", "-c", "cd \""+dir+"\"; python3 \""+path+"\"; exec bash") // Wszytskie systemy inne niż Windows i MacOS
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "start cmd.exe /K python \""+path+"\"")
    } else if runtime.GOOS == "darwin" {
        cmd = exec.Command("echo", macos) // Jeśli umiesz, to możesz se to poprawić i mi powiedzieć jak to zrobić, z góry dzięki
        dialog.ShowInformation("Linux >> Windows >> MacOS", macos, w)
	}
	err := cmd.Start()
	if err != nil {
		log.Println(erro, err)
	}
}

func openTerminal(w fyne.Window) {
	if openFolderPath == "" {
		dialog.ShowInformation("Błąd", "Najpierw wybierz folder projektu!", w)
		return
	}

	var cmd *exec.Cmd
	dir := openFolderPath

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "cmd.exe", "/K", "cd", "\"", dir, "\"")
	} else if runtime.GOOS == "darwin" {
	    dialog.ShowInformation("Linux >> Windows >> MacOS", macos, w)
        cmd = exec.Command("echo", macos) // Jeśli umiesz, to możesz se to poprawić i mi powiedzieć jak to zrobić, z góry dzięki
	} else {
		cmd = exec.Command("x-terminal-emulator", "-e", "bash", "-c", "cd \""+dir+"\"; exec bash") // Wszytskie systemy inne niż Windows i MacOS
	}
	err := cmd.Start()
	if err != nil {
		log.Printf(erro, err)
	}
}
