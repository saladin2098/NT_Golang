package main

import "time"

func Date(date chan time.Time) {
	chan1 := make(chan time.Time)
	chan1 <- time.Date(2024,04,23,16,3,0,0,time.UTC)
	go main(chan1)
}
func main(date chan time.Time) {
	
}
