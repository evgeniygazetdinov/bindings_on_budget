package main

import (
	"log"
	"time"
)

// Run the function every tick
// Return false from the func to stop the ticker
func Every(duration time.Duration, work func(time.Time) bool) chan bool {
	ticker := time.NewTicker(duration)
	stop := make(chan bool, 1)

	go func() {
		defer log.Println("ticker stopped")
		for {
			select {
			case time := <-ticker.C:
				if !work(time) {
					stop <- true
				}
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func f() func() int {
	i := 0

	return func() int {
		i++
		return i
	}

}
func main() {

	// fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
	stop := Every(1*time.Second, func(time.Time) bool {

		log.Println("tick")

		return true
	})
	time.Sleep(10 * time.Second)
	stop <- true

	// time.Sleep(3 * time.Second)
}
