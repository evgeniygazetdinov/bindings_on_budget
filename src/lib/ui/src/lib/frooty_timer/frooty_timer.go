package frooty_timer

import (
    "fmt"
    "time"
    "github.com/therecipe/qt/widgets"
  // notif "./src/lib/notificator"	
   "sync"
   // "math"
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

func workLoop(begin int, realBegin int, runnerChannel chan int, timeLabel *widgets.QLabel){
    var wg sync.WaitGroup
    my_chan := make(chan string)
    defaultTime := begin
    for i:= 0; i<begin; i++{
        select{
        case value, ok := <-runnerChannel:
            if ok {
                if value > 0{
                  // notif.NOTIFY_ME()
                    fmt.Println("values is 0")
                    wg.Wait()
                    workLoop(begin, defaultTime, runnerChannel, timeLabel)
                }
            } else {
                fmt.Println("Channel closed!")
            }     
        default:      
            // if timebegin < 0{
            //     wg.Wait()
            //     return 
            // }
            go getChars(my_chan, begin)
            wg.Add(1)
            realBegin-=1
            timeLabel.SetText(<-my_chan) 
            time.Sleep(1 * time.Second)
            begin-=1;        
        }
    }

}

func FROOTY_TIMER(timeBegin int, placeForUse *widgets.QLabel, runFlag chan int){
    begin := calculateLimitOfTime(timeBegin) 
    fmt.Println(begin)
    workLoop(begin, timeBegin, runFlag, placeForUse)    
}