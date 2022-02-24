package frooty_timer

import (
    "fmt"
    "time"
    "github.com/therecipe/qt/widgets"
  // notif "./src/lib/notificator"	
   //"sync"
)

func getChars(cs chan string, begin int ) {
    minutes := (begin % 3600) / 60;
    seconds := begin % 60;
    timeString := fmt.Sprintf("%02d:%02d", minutes, seconds);
    cs <- timeString
    fmt.Println("get chars over")
}

func FROOTY_TIMER(timebegin int, place_for_use *widgets.QLabel, run_flag chan int){
    my_chan := make(chan string)
    begin := timebegin
    for i:= 0; i<begin; i++{
        select {
        case value, ok := <-run_flag:
            if ok {
                if value > 0{
                  // notif.NOTIFY_ME()
                    fmt.Println("values is 0")
                    return
                }
            } else {
                fmt.Println("Channel closed!")
            }
        default:
            go getChars(my_chan, begin)
            //wait group herer
            place_for_use.SetText(<-my_chan) 
            time.Sleep(1 * time.Second)
            begin-=1;
                 
        }
   
      
    }
    
}