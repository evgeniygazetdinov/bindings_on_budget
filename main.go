package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
	//"reflect"
	ui "evgeniygazetdinov/bindings_on_budget/src/lib/ui"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	ui.UI_SHIT(window)
	window.Show()
	app.Exec()
}
