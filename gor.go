package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func getChars(cs chan string, begin int ) {
    
    minutes := (begin % 3600) / 60;
    seconds := begin % 60;
    timeString := fmt.Sprintf("%02d:%02d", minutes, seconds);
    cs <- timeString

}

func main() {
    fmt.Println("main execution started at time", time.Since(start))
    my_chan := make(chan string)
    begin := 1500
    for i:= 0; i<begin; i++ {
        go getChars(my_chan, begin)
        begin-=1;
        time.Sleep(1 * time.Second)
        fmt.Println(<-my_chan)
    }
    
    

    fmt.Println("\nmain execution stopped at time", time.Since(start))
}