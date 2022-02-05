package main

import (
    "fmt"
    "time"
    "math/rand"
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

func mY_counter(rutine_name string){
    
    for i:=0;i<1000;i++{
        r := rand.Intn(4)
        time.Sleep(time.Duration(r) * time.Second)
        fmt.Println("whis is rutine is", rutine_name)
        
    }
}

func main() {
   go mY_counter("one")
   go mY_counter("two")
   go mY_counter("three")
   go mY_counter("four")
   FrootyTimer(1500)
   
}