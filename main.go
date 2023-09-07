package main

import (
	v35 "gosetgo/v3.5"
	"log"
	"sync"
)

func main() {

	// add 10 items to the set
	n := 10

	var wg sync.WaitGroup

	// add 10 items to the set
	s := v35.NewSet()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			added := s.Add(i)
			log.Println(s.Contains(i - 1))
			log.Println("Added", i, "to the set?", added)
		}(i)
	}

	wg.Wait()
}
