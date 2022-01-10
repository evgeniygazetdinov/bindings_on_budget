package main

import (
	"fmt"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"os"
	//"reflect"
	"strconv"
)

func equalize_operation(operator string, first_digit int, second_digit int) int {
	var result int
	switch operator {
	case "+":
		result = first_digit + second_digit
	case "-":
		result = first_digit - second_digit
	}
	return result
}

func operation(label *widgets.QLabel, input_text string, operator string) {
	res := label.Text()
	res_int, err := strconv.Atoi(res)
	add_integer, err := strconv.Atoi(input_text)
	fmt.Println(res_int, add_integer)
	res = strconv.Itoa(equalize_operation(operator, res_int, add_integer))
	if err != nil {
		fmt.Println("error")
	}
	label.SetText(res)
}

func operation_input(label *widgets.QLabel, input *widgets.QLineEdit, operator string){		
	if label.Text() == "0" {
		label.SetText(input.Text())
	} else {
		operation(label, input.Text(), operator)
	}

}

func ui_shit(window *widgets.QMainWindow){
	window.SetWindowTitle("budget")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	//input.SetValidator(widgets.NewQRegExpValidator(widgets.NewQRegExp("[0-9]*"), &input));
	input.SetValidator(gui.NewQIntValidator(input))
	input.SetPlaceholderText("")
	layout.AddWidget(input, 1, 0)

	// Create a button and add it to theco layout
	plus := widgets.NewQPushButton2("+", nil)
	minus := widgets.NewQPushButton2("-", nil)
	layout.AddWidget(plus, 2, 0)
	layout.AddWidget(minus, 3, 0)
	label := widgets.NewQLabel2("0", nil, 0)
	layout.AddWidget(label, 0, 0)

	// Connect event for button
	plus.ConnectClicked(func(checked bool) {
		operation_input(label, input, "+")
	})
	minus.ConnectClicked(func(checked bool) {
		operation_input(label, input, "-")
	})


	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)
}

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	ui_shit(window)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}
