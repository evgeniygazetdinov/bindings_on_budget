package main

import (
	ui "evgeniygazetdinov/bindings_on_budget/src/lib/ui"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main22() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	ui.UI_SHIT(window)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}
