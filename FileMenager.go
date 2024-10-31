package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func loadFolder(editor *widget.Entry, dir string, accordion *widget.Accordion, w fyne.Window) {
	items := make(map[string][]string)
	openFolderPath = dir
	ActFIles(items, dir)

	accordion.CloseAll()

	for folder, files := range items {
		content := container.NewVBox()
		for _, file := range files {
			filePath := file

			if strings.HasSuffix(filePath, ".MaqPyE-Backup") { // Pomijanie plików backupów (nie dodaje ich do Acordiona)
				continue
			}

			button := widget.NewButton(filepath.Base(filePath), func() {
				ext := filepath.Ext(filePath)
				for _, disallowed := range NoExt {
					if ext == disallowed {
						dialog.ShowConfirm("Unsupported File!",
							unsfile,
							func(confirmed bool) {
								if confirmed {
									openFile(editor, filePath)
								}
							}, w)
						return
					}
				}

				fileInfo, err := os.Stat(filePath)
				if err != nil {
					log.Println(erro, err)
					return
				}
				if fileInfo.Size() > mb {
					dialog.ShowConfirm("Large File!",
						lgrFile,
						func(confirmed bool) {
							if confirmed {
								openFile(editor, filePath)
							}
						}, w)
					return
				}
				openFile(editor, filePath)
			})
			content.Add(button)
		}
		accordion.Append(widget.NewAccordionItem(folder, content))
	}
}

func openFile(editor *widget.Entry, filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	editor.SetText(string(content))
	currentFilePath = filePath
}

func ActFIles(items map[string][]string, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(erro, err)
		return
	}

	for _, file := range files {
		itemPath := filepath.Join(path, file.Name())
		if file.IsDir() {
			if _, exists := items[file.Name()]; !exists {
				items[file.Name()] = []string{}
			}
			ActFIles(items, itemPath)
		} else {
			parentDir := filepath.Base(path)
			items[parentDir] = append(items[parentDir], itemPath)
		}
	}
}
