package ui

import (
	// "github.com/therecipe/qt/gui"
	//"github.com/therecipe/qt/core"

	timer "evgeniygazetdinov/bindings_on_budget/src/lib/ui/src/lib/frooty_timer"

	"github.com/therecipe/qt/widgets"
	// cal "./src/lib/cal"
	// db "evgeniygazetdinov/bindings_on_budget/src/lib/db"
)

const DURATION = 1200

func update_gui(label *widgets.QLabel, startButton *widgets.QPushButton) {
	timer.FROOTY_TIMER(DURATION, label, startButton)
}

func UI_SHIT(window *widgets.QMainWindow) {
	window.SetWindowTitle("budget")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)
	input := widgets.NewQLineEdit(nil)
	//input.SetValidator(gui.NewQIntValidator(input))
	input.SetPlaceholderText("")
	layout.AddWidget(input, 1, 0)
	addButton := widgets.NewQPushButton2("добавить задачу", nil)
	layout.AddWidget(addButton, 2, 0)
	// Create a button and add it to theco layout
	currentTask := widgets.NewQLabel2(" ", nil, 0)
	layout.AddWidget(currentTask, 0, 0)

	timeLabel := widgets.NewQLabel2("0", nil, 0)
	layout.AddWidget(timeLabel, 0, 0)
	// Connect event for button
	startButton := widgets.NewQPushButton2("start", nil)
	layout.AddWidget(startButton, 2, 0)
	// handler start
	isStartPushed := true
	// fmt.Println(reflect.TypeOf(isStartPushed))
	startButton.ConnectClicked(func(checked bool) {
		if isStartPushed {
			isStartPushed = false
			update_gui(timeLabel, startButton)
		}

	})

	// currentTask.SetText(db.GET_LAST_TASK())

	// addButton.ConnectClicked(func(checked bool) {
	// 	if input.Text() != " " && !db.TASK_EXIST_IN_DB(input.Text()) {
	// 		db.QUERY_INPLACE(fmt.Sprintf("insert into tasks(tasks_id, name) values((SELECT MAX(tasks_id)from tasks)+1,'%s');", input.Text()))
	// 		currentTask.SetText(input.Text())
	// 	}
	// })

	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)
}
