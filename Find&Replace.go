package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"strings"
)

// Find & Replace
func showFindReplaceDialog(editor *widget.Entry, w fyne.Window) {
	findEntry := widget.NewEntry()
	findEntry.SetPlaceHolder("Szukany tekst...")
	
	replaceEntry := widget.NewEntry()
	replaceEntry.SetPlaceHolder("Zamień na...")

	// Zmiana jednego
	replaceOne := widget.NewButton("Zamień", func() {
		Ft := findEntry.Text
		Rt := replaceEntry.Text
		content := editor.Text
		
		if Ft == "" {
			dialog.ShowInformation("Błąd", wprowtxtdosz, w)
			return
		}
		
		if strings.Contains(content, Ft) {
			newContent := strings.Replace(content, Ft, Rt, 1)
			editor.SetText(newContent)
		} else {
			dialog.ShowInformation("Info", nznszktx, w)
		}
	})
	
	// ALL
	replaceAll := widget.NewButton("Zamień wszystko", func() {
		Ft := findEntry.Text
		Rt := replaceEntry.Text
		content := editor.Text
		
		if Ft == "" {
			dialog.ShowInformation("Błąd", wprowtxtdosz, w)
			return
		}
		
		// Zm All
		newContent := strings.ReplaceAll(content, Ft, Rt)
		editor.SetText(newContent)
		dialog.ShowInformation("Sukces", "Zamieniono wszystkie wystąpienia!", w)
	})

	// Okno
	dialog.ShowCustom("Find & Replace", "Zamknij", container.NewVBox(findEntry, replaceEntry, replaceOne, replaceAll), w)
}
