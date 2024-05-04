package main

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}

	wg.Done()

	// return
}

func main() {
	s1 := 0
	log.Println(s1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 1000000000, wg, &s1)
	wg.Wait()

	log.Println(s1)

}
