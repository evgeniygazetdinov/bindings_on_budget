package main

import (

	"github.com/therecipe/qt/widgets"
	"os"
	//"reflect"
	ui "./src/lib/ui"
)


// func ui_shit(window *widgets.QMainWindow){
// 	window.SetWindowTitle("budget")
// 	window.SetMinimumSize2(200, 200)

// 	// Create main layout
// 	layout := widgets.NewQVBoxLayout()

// 	// Create main widget and set the layout
// 	mainWidget := widgets.NewQWidget(nil, 0)
// 	mainWidget.SetLayout(layout)

// 	// Create a line edit and add it to the layout
// 	input := widgets.NewQLineEdit(nil)
// 	input.SetValidator(gui.NewQIntValidator(input))
// 	input.SetPlaceholderText("")
// 	layout.AddWidget(input, 1, 0)

// 	// Create a button and add it to theco layout
// 	plus := widgets.NewQPushButton2("+", nil)
// 	minus := widgets.NewQPushButton2("-", nil)
// 	layout.AddWidget(plus, 2, 0)
// 	layout.AddWidget(minus, 3, 0)
// 	label := widgets.NewQLabel2("0", nil, 0)

// 	layout.AddWidget(label, 0, 0)
// 	// Connect event for button
// 	plus.ConnectClicked(func(checked bool) {
// 		cal.OPERATION_INPUT(label, input, "+")
// 	})
// 	minus.ConnectClicked(func(checked bool) {
// 		cal.OPERATION_INPUT(label, input, "-")
// 		notif.NOTIFY_ME()
// 	})


// 	// Set main widget as the central widget of the window
// 	window.SetCentralWidget(mainWidget)
// }

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	// ui_shit(window)
	ui.UI_SHIT(window)
	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}
