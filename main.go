package main

import (

	"github.com/therecipe/qt/widgets"
	"os"
	//"reflect"
	ui "./src/lib/ui"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	ui.UI_SHIT(window)
	window.Show()
	app.Exec()
}
