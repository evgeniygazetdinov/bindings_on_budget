package main

import (
	"fmt"
	"github.com/therecipe/qt/widgets"
	"os"
	"reflect"
	"strconv"
)

func operation(label *widgets.QLabel, input_text string) {
	res := label.Text()
	res_int, err := strconv.Atoi(res)
	add_integer, err := strconv.Atoi(input_text)
	fmt.Println(res_int, add_integer)
	res = strconv.Itoa(res_int + add_integer)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(reflect.TypeOf(add_integer), reflect.TypeOf(res_int))
	label.SetText(res)
}

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	//input.SetValidator(widgets.NewQRegExpValidator(widgets.NewQRegExp("[0-9]*"), &input));
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 1, 0)

	// Create a button and add it to the layout
	button := widgets.NewQPushButton2("2. click me", nil)
	layout.AddWidget(button, 2, 0)
	label := widgets.NewQLabel2("0", nil, 0)
	layout.AddWidget(label, 0, 0)

	// Connect event for button
	button.ConnectClicked(func(checked bool) {
		//widgets.QMessageBox_Information(nil, "OK", input.Text(),
		//widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		fmt.Println(label.Text() == "")

		if label.Text() == "0" {

			label.SetText(input.Text())
		} else {
			operation(label, input.Text())
		}
	})

	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}
