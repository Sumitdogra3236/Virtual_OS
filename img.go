package main

import (
	"fmt"
    "io/ioutil"
    "log"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	root_src := "C:\\Users\\sumit\\OneDrive\\Pictures\\Screenshots";
	files, err := ioutil.ReadDir(root_src);
	if err != nil {
        log.Fatal(err)
    }
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
    }


	w.ShowAndRun()
}

