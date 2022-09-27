package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race Condition - LCO")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("One R")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Two R")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("Three R")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.RLock()
		fmt.Println("Four R", score)
 		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println("Score", score)

}
