package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"strconv"
)

// --------------------------------- DATA ------------------------------------------ \\


var macos string = "\"Siema, ogólnie to nie znam się na MacOS i nie wiem jak tu to mam zainplementować, więc jak coś, to sorry ale ta funkcja u ciebie nie zadziała\""
var unsfile string = "Plik tego typu może być duży i powodować problemy przy otwieraniu. Czy na pewno chcesz go otworzyć?"
var lgrFile string = "Plik jest większy niż 1 MB. Czy na pewno chcesz go otworzyć?"
var wprowtxtdosz string = "Wprowadź tekst do wyszukania" // serio przepraszam za te nazwy XDDD
var erro string = "\x1b[31;1mError!\x1b[0m, %s..."
var nznszktx string = "Nie znaleziono szukanego tekstu"

var NoExt = []string{".png", ".mp4", ".ogg", ".mp3", ".webp", ".webm"} // "Złe" rozszerzenia, możesz se tu coś dodać innego jak chcesz

var mb int64 = 1024*1024

var currentFilePath string
var openFolderPath string

var w fyne.Window


// --------------------------------- CODE ------------------------------------------ \\

func main() {
	a := app.New()
	w := a.NewWindow("MaqPyE")
	w.Resize(fyne.NewSize(800, 600))

	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Witaj w BECIE MaqPyE! Dobra pisz coś...")

	// LineCOUNT!
	lineNumbers := widget.NewLabel("")
	updateLineNumbers := func() {
		lineCount := strings.Count(editor.Text, "\n") + 1
		lines := ""
		for i := 1; i <= lineCount; i++ {
			lines += strconv.Itoa(i) + "\n"
		}
		lineNumbers.SetText(lines)
	}

	lineNumbersScroll := container.NewVScroll(lineNumbers)
	lineNumbersScroll.SetMinSize(fyne.NewSize(25, 0)) // Zmiejsz to sobie jak chcesz, to jest szerokość paska z numeracją linii (btw jeśli to czytasz to znaczy że przeszłem na dobrą ścieżke i projket jest OpenSource, JEJJEJEJ!!!!)
	editorScroll := container.NewVScroll(editor)

    editorScroll.OnScrolled = func(_ fyne.Position) {
        updateLineNumbers()
    }

	editor.OnChanged = func(text string) { // Aktualizacja Lini
		updateLineNumbers()
	}

    // tu masz główne layouty debilu
	editorContainer := container.NewHSplit(lineNumbersScroll, editorScroll)
	editorContainer.Offset = 0.05

	accordion := widget.NewAccordion()
	mainContainer := container.NewHSplit(accordion, editorContainer)

    terminalData, err := ioutil.ReadFile("Icons/Terminal.png")
    if err != nil {
        log.Fatal(err)
    }
    terminalIcon := fyne.NewStaticResource("terminal", terminalData)

	runButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		if currentFilePath != "" {
			runPythonScript(currentFilePath)
		} else {
			dialog.ShowInformation("Info", "Nie wybrano pliku do uruchomienia.", w)
		}
	})
	terminalButton := widget.NewButton("", func() {
		openTerminal(w)
	})
	terminalButton.SetIcon(terminalIcon)
	settingsButton := widget.NewButton("Ustawienia", func() {
		EasterEgg(w)
	})

	// Find & Replace
	findReplaceButton := widget.NewButton("Find & Replace", func() {
		showFindReplaceDialog(editor, w)
	})

	// ToolBar
	toolbar := container.NewBorder(nil, nil, container.NewHBox(runButton, terminalButton, findReplaceButton), settingsButton)

	// LayOuty
	layout := container.NewBorder(nil, toolbar, nil, nil, mainContainer)
	w.SetContent(layout)

	// Otwieranie workcpace
	openFolderAction := widget.NewButton("Open Folder", func() {
		dialog.NewFolderOpen(func(dir fyne.ListableURI, err error) {
			if err == nil && dir != nil {
				loadFolder(editor, dir.Path(), accordion, w)
			}
		}, w).Show()
	})

	saveFile := func() {
		if currentFilePath != "" {
			err := ioutil.WriteFile(currentFilePath, []byte(editor.Text), 0644)
			if err != nil {
				log.Println("Error saving file:", err)
			}
		}
	}

	// dodanie menu
	mn := fyne.NewMenu("File",
		fyne.NewMenuItem("Open Folder", func() {
			openFolderAction.OnTapped()
		}),
		fyne.NewMenuItem("Save", func() {
			saveFile()
		}),
	)
	w.SetMainMenu(fyne.NewMainMenu(mn))

	// Ctrl + S, który jak zawsze nie działa
	w.Canvas().AddShortcut(&fyne.ShortcutPaste{}, func(_ fyne.Shortcut) {
		saveFile()
	})
	
	// Ctrl+F dla Find & Replace (też nie działa)
	w.Canvas().AddShortcut(&fyne.ShortcutCopy{}, func(_ fyne.Shortcut) {
		showFindReplaceDialog(editor, w)
	})

	// Backup co minutę
	go func() {
		for range time.Tick(1 * time.Minute) {
			if currentFilePath != "" {
				backupPath := currentFilePath + ".MaqPyE-Backup"
				err := ioutil.WriteFile(backupPath, []byte(editor.Text), 0644)
				if err != nil {
					log.Println(erro, err)
				}
			}
		}
	}()

	w.ShowAndRun()
}

func EasterEgg(w fyne.Window) {
	img := canvas.NewImageFromFile("Icons/cos.jpg")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(480, 270))
	label := widget.NewLabel("Nie zdąrzyłem jeszcze zrobić ustawień, ale jako iż to konkurs Hallowenowy... masz EasterEgga XD")

	content := container.NewVBox(img, label)

	dialog := dialog.NewCustom("EasterEgg", "XD", content, w)
	dialog.Show()
}
