package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type menuItem struct {
	label   string
	command string
}

func runCommand(item menuItem) func() {
	return func() {
		parts := strings.Fields(item.command)
		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}

func main() {

	// if the number of args isn't divisible by 2, we done goofed
	if (len(os.Args)-1)%2 != 0 {
		fmt.Fprintln(os.Stderr, "error: Invalid number of arguments")
		os.Exit(1)
	}

	// make menu items from the args
	var menuItems []menuItem
	for i := 1; i < len(os.Args); i += 2 {
		label := os.Args[i]
		command := os.Args[i+1]
		menuItems = append(menuItems, menuItem{label, command})
	}

	// hello andy :)
	a := app.New()
	w := a.NewWindow("8menu")

	// add the buttons and their corresponding functions
	buttons := make([]fyne.CanvasObject, len(menuItems)+1)
	for i, item := range menuItems {
		buttons[i] = widget.NewButton(item.label, runCommand(item))
	}

	// add a close button
	exitButton := widget.NewButton("Exit", func() {
		w.Close()
	})
	buttons[len(menuItems)] = exitButton

	// make the layout
	content := container.NewVBox(buttons...)
	w.SetContent(content)

	// set min width
	minHeight := 1
	w.Resize(fyne.NewSize(150, float32(minHeight)))

	w.ShowAndRun()
}
