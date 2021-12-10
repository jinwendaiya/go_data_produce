package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Microsecond * 1000)
	//defer ticker.Stop() // 【重要】关闭ticker
	done := make(chan bool)
	time.Sleep(2 * time.Second)
	go func() {
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}






}
