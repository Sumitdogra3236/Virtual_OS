// package main

// import("fmt") 

// func main(){
// 	fmt.Print("Hello to GO Programming Language")
// }

package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	historyBtn := widget.NewButton("History", func(){

	})

	backBtn := widget.NewButton("Back",func() {

	})

	clearBtn := widget.NewButton("Clear", func() {

	})

	openBtn := widget.NewButton("Clear", func() {

	})

	closeBtn := widget.NewButton("Clear", func() {

	})

	divideBtn := widget.NewButton("Clear", func() {

	})

	sevenBtn := widget.NewButton("Clear", func() {

	})

	eightBtn := widget.NewButton("Clear", func() {

	})

	nineBtn := widget.NewButton("Clear", func() {

	})

	multiplyBtn := widget.NewButton("Clear", func() {

	})

	fourBtn := widget.NewButton("Clear", func() {

	})

	fiveBtn := widget.NewButton("Clear", func() {

	})

	sixBtn := widget.NewButton("Clear", func() {

	})

	minusBtn := widget.NewButton("Clear", func() {

	})

	oneBtn := widget.NewButton("Clear", func() {

	})

	twoBtn := widget.NewButton("Clear", func() {

	})

	threeBtn := widget.NewButton("Clear", func() {

	})
	plusBtn := widget.NewButton("Clear", func() {

	})

	zeroBtn := widget.NewButton("Clear", func() {

	})

	dotBtn := widget.NewButton("Clear", func() {

	})

	equalBtn := widget.NewButton("Clear", func() {

	})
	

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		container.NewGridWithColumns(1,
		container.NewGridWithColumns(2, 
			historyBtn,
			backBtn,
			),

			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
			nineBtn,
			eightBtn,
			sevenBtn,
			multiplyBtn),

			container.NewGridWithColumns(4,
			fourBtn,
			fiveBtn,
			fiveBtn,
			sixBtn,
			minusBtn),
		),
		container.NewGridWithColumns(4,
		oneBtn,
		twoBtn,
		threeBtn,
		plusBtn),
		container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
		zeroBtn,
		dotBtn),
		eightBtn,
	),
	))

	w.ShowAndRun()
}
