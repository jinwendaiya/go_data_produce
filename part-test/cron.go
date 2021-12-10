package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1ms", func() {
		fmt.Println("tick every 1 Millisecond")
	})
	//time.ParseDuration()
	c.Start()
	//time.Sleep(time.Second * 5)

	for {
		time.Sleep(1*time.Millisecond)
	}

}
