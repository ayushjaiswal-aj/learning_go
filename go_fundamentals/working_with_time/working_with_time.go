package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	currentTime := time.Now()
	pl("Current time:", currentTime)
	pl("Current time:", currentTime.Format("2006-01-02 15:04:05"))
	pl(currentTime.Year(), ":", currentTime.Month(), ":", currentTime.Day(), " ", currentTime.Hour(), ":", currentTime.Minute(), ":", currentTime.Second())
}
