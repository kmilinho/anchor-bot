package cli

import (
	"fmt"
	"os"
)

func Run() {
	fmt.Println("I'm running!")
	events := make(chan string)
	stop := make(chan int8)
	go commandHandler(events, stop)
	<-events
}
