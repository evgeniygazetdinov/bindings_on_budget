package main

import (
    "fmt"
    "time"
)

var start time.Time

func init() {
    start = time.Now()
}

func getChars(cs chan string) {
    begin := 1500
    for i:= 0; i<begin; i++ {
        minutes := (begin % 3600) / 60;
        seconds := begin % 60;
        timeString := fmt.Sprintf("%02d:%02d", minutes, seconds);
        begin-=1;
        time.Sleep(1 * time.Second)
        cs <- timeString
    }
}

func main() {
    fmt.Println("main execution started at time", time.Since(start))
    my_chan := make(chan string)
    go getChars(my_chan)
    time.Sleep(10 * time.Second)
    fmt.Println(<-my_chan)

    fmt.Println("\nmain execution stopped at time", time.Since(start))
}