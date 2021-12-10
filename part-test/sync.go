package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(2)

	var mut sync.Mutex

	wg.Add(5)

	go func() {
		mut.Lock()
		defer mut.Unlock()
		fmt.Println(1)
		wg.Done()
	}()

	go func() {
		fmt.Println(2)
		wg.Done()
	}()
	go func() {
		fmt.Println(3)
		wg.Done()
	}()
	go func() {
		fmt.Println(4)
		wg.Done()
	}()
	go func() {
		fmt.Println(5)
		wg.Done()
	}()

	//runtime



	wg.Wait()
}