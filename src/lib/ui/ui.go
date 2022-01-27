package ui

import (
	// "github.com/therecipe/qt/gui"
	//"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/honeycombio/beeline-go/timer"
	// "reflect"
	"fmt"
	// cal "./src/lib/cal"
	// notif "./src/lib/notificator"
)
const DURATION = 1500 


func update_gui(label *widgets.QLabel){
	label.SetText("1")
}

func UI_SHIT(window *widgets.QMainWindow){
	window.SetWindowTitle("budget")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	// input := widgets.NewQLineEdit(nil)
	// input.SetValidator(gui.NewQIntValidator(input))
	// input.SetPlaceholderText("")
	// layout.AddWidget(input, 1, 0)

	// Create a button and add it to theco layout
	time_label := widgets.NewQLabel2("0", nil, 0)
	layout.AddWidget(time_label, 0, 0)
	// Connect event for button
	startButton := widgets.NewQPushButton2("start", nil)
	stopButton := widgets.NewQPushButton2("stop", nil)
	layout.AddWidget(startButton, 2, 0)
	layout.AddWidget(stopButton, 2, 0)
	startButton.ConnectClicked(func(checked bool) {
		timer := timer.Start()
		// timer := core.NewQBasicTimer()
		// fooType := reflect.TypeOf(timer)
		// for i := 0; i < fooType.NumMethod(); i++ {
		// method := fooType.Method(i)
		// fmt.Println(method.Name)
		// }
		// update_gui(time_label)
		// timer.Start(DURATION)
		dur := timer.Finish()
		fmt.Printf("%g", dur)
		update_gui(time_label)
	})


	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)
}

