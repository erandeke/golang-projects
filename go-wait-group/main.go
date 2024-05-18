package main

import (
	"fmt"
	"sync"
)

func main() {

	//wait group waits for the collection of go routines to be finished
	//The main goroutine adds to the set the number of gorutines to wait for

	var sleeper sync.WaitGroup
	ninjas_names := []string{"Tommy", "Harry", "Dick"}
	sleeper.Add(len(ninjas_names))
	for _, evil := range ninjas_names {
		go attackNinja(evil, &sleeper)
	}
	sleeper.Wait()
	fmt.Println("Mission completed")

}

func attackNinja(target string, sleeper *sync.WaitGroup) {
	fmt.Println("Attacked by", target)
	sleeper.Done()

}
