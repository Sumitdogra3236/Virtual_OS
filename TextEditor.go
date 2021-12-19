package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var count int = 1
func main() {
	a := app.New()

	w := a.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(600,600))

	content := container.NewVBox(
	container.NewHBox(
		widget.NewLabel("Text Editor"),
	),
)
	content.Add(widget.NewButton("Add New File", func(){
	content.Add(widget.NewLabel("New File"+strconv.Itoa(count)))
	count++
}))


	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	input.Resize(fyne.NewSize(1200,1200))


	saveBtn := widget.NewButton("Save text File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error){
				textData := []byte(input.Text)

				uc.Write(textData)
			},w)

			saveFileDialog.SetFileName("New File"+strconv.Itoa(count-1) +".txt")

			saveFileDialog.Show()
	})


	
	w.SetContent(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
			saveBtn,
			),
		),
		
	)
	w.ShowAndRun()
}