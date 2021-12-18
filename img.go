package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800,600))

	root_src := "C:\\Users\\sumit\\OneDrive\\Pictures\\Screenshots";
	files, err := ioutil.ReadDir(root_src);
	if err != nil {
        log.Fatal(err)
    }
	var picsArr[] string; 
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")[1];
			if extension == "png" || extension == "jpeg"{
			picsArr = append(picsArr,root_src+"\\"+file.Name())
			}
		}
    }

	image := canvas.NewImageFromFile(picsArr[0])
	tabs := container.NewAppTabs(
		container.NewTabItem("Image1",)
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	w.SetContent(container.NewHBox(image,tabs))
	w.ShowAndRun()
}

