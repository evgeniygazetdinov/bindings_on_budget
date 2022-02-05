package main

import (
    "fmt"
    "time"
)

func getChars(cs chan string, begin int ) {
    minutes := (begin % 3600) / 60;
    seconds := begin % 60;
    timeString := fmt.Sprintf("%02d:%02d", minutes, seconds);
    cs <- timeString

}

func FrootyTimer(timebegin int){
    my_chan := make(chan string)
    begin := timebegin
    for i:= 0; i<begin; i++ {
        go getChars(my_chan, begin)
        begin-=1;
        time.Sleep(1 * time.Second)
        fmt.Println(<-my_chan)
    }
    
}

func main() {
    FrootyTimer(1500)
}