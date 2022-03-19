package main

import (
	ui "./src/lib/ui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
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
