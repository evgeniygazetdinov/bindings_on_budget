package frooty_timer

import (
    "fmt"
    "time"
    "github.com/therecipe/qt/widgets"
  // notif "./src/lib/notificator"	
//   "sync"
   // "math"
//    "reflect"
)

func calculateLimitOfTime(beforeBegiTime int) int{
    return (beforeBegiTime * 2) + int(beforeBegiTime / 2)
}

func getChars(cs chan string, begin int ) {
    minutes := (begin % 3600) / 60;
    seconds := begin % 60;
    timeString := fmt.Sprintf("%02d:%02d", minutes, seconds);
    cs <- timeString
    fmt.Println("get chars over")
    return 
}

func workLoop(begin int, realBegin int, startButton *widgets.QPushButton,
     timeLabel *widgets.QLabel){
    my_chan := make(chan string)
    defaultTime := realBegin
    for i:= 0; i<begin; i++{
        // select{
        // case value, ok := <-runnerChannel:
        //     if ok {
        //         if value < 0{
        //             begin = defaultTime
        //             runnerChannel <- 1
        //             return 
        // //         //    wg.Wait()
        //         //    workLoop(begin, defaultTime, runnerChannel, timeLabel)
        //         }
        //     } else {
        //         fmt.Println("Channel closed!")
        //     }     
        // default:      
            // if realBegin == 0{
            //     close(my_chan)
            //     return 
            // }
            // fmt.Println(realBegin)
    
            fmt.Println("begin")
            go getChars(my_chan, realBegin)
            fmt.Println("set to label")
            timeLabel.SetText(<-my_chan)
            fmt.Println("setted") 
            time.Sleep(1 * time.Second)
            fmt.Println("after delay")
            begin-=1;
            realBegin-=1
            fmt.Println("after minus ")
            fmt.Println(i)
            if i == defaultTime{
                startButton.SetDisabled(false)
            } 

                
    }

}


func FROOTY_TIMER(timeBegin int, placeForUse *widgets.QLabel, startButton *widgets.QPushButton){
    begin := calculateLimitOfTime(timeBegin)
    //fmt.Println(runFlag)
    workLoop(begin, timeBegin, startButton,
             placeForUse)    
}