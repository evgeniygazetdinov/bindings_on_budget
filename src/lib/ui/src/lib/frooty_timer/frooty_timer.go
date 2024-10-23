package frooty_timer

import (
	"fmt"
	"log"
	"time"

	"github.com/therecipe/qt/widgets"
	// notif "./src/lib/notificator"
	//   "sync"
	// "math"
	//    "reflect"
)

func calculateLimitOfTime(beforeBegiTime int) int {
	return (beforeBegiTime * 2) + int(beforeBegiTime/2)
}

func getChars(cs chan string, begin int) {
	minutes := (begin % 3600) / 60
	seconds := begin % 60
	timeString := fmt.Sprintf("%02d:%02d", minutes, seconds)
	cs <- timeString
	return
}

func Every(duration time.Duration, work func(time time.Time, variable chan bool, updateFunc func()) bool) chan bool {
	// ticker := time.NewTicker(duration)
	stop := make(chan bool, 1)

	go func() {
		defer log.Println("ticker stopped")
		for {
			select {
			// case time := <-ticker.C:
			// 	// if !work(time, variable, updateFunc) {

			// 	// 	stop <- true
			// 	// }
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func UpdateTimeLabel(begin, realBegin int, timeLabel *widgets.QLabel) {
	my_chan := make(chan string)
	go getChars(my_chan, realBegin)
	timeLabel.SetText(<-my_chan)
	begin -= 1
	realBegin -= 1
}

func workLoop(begin int, realBegin int, startButton *widgets.QPushButton,
	timeLabel *widgets.QLabel,
	chanPush chan bool) {

	Every(1*time.Second, func(t time.Time, v chan bool, UpdateLabel func()) bool {
		my_chan := make(chan string)
		go getChars(my_chan, realBegin)
		timeLabel.SetText(<-my_chan)
		begin -= 1
		realBegin -= 1
		return true
	})
	fmt.Printf("%v heres", <-chanPush)
}

func FROOTY_TIMER(timeBegin int, placeForUse *widgets.QLabel,
	startButton *widgets.QPushButton,
	chanPush chan bool) {
	begin := calculateLimitOfTime(timeBegin)
	workLoop(begin, timeBegin, startButton,
		placeForUse, chanPush)

}
