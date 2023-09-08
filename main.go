package main

import (
	v6 "gosetgo/v6"
	"log"
	"sync"
)

func main() {

	// add 10 items to the set
	n := 10

	var wg sync.WaitGroup

	// add 10 items to the set
	s := v6.NewSet()
	for i := 1; i < n+1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			added := s.Add(i)
			log.Println(s.Contains(i - 1))
			log.Println("Added", i, "to the set?", added)
		}(i)
	}

	wg.Wait()

	// remove 5 items from the set
	for i := 1; i < n/2+1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			removed := s.Remove(i)
			log.Println("Removed", i, "from the set?", removed)
		}(i)
	}

	wg.Wait()

	// check if 5 items are in the set
	for i := 1; i < n/2+1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println("Contains", i, "in the set?", s.Contains(i))
		}(i)
	}

	wg.Wait()

	// check if other 5 items are in the set
	for i := n / 2; i < n+1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println("Contains", i, "in the set?", s.Contains(i))
		}(i)
	}

	wg.Wait()
}
