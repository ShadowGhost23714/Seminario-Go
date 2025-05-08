package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	paper = iota
	grass
	match
)

var smokers = map[int]string{
	paper: "Sandy",
	grass: "Apple",
	match: "Daisy",
}

var wg sync.WaitGroup

func arbitrate(signals, ingredients [3]chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		next := rand.Intn(3)
		fmt.Println("\nNext:", smokers[next])
		signals[next] <- next
		/* for c := range ingredients {
			if c != next {
				ingredients[c] <- 1
			}
		} */
		for j := 0; j < 2; j++ {
			select {
			case ingredients[paper] <- 1:
			case ingredients[grass] <- 1:
			case ingredients[match] <- 1:
			}
		}
	}
	for c := range signals {
		close(signals[c])
	}
	wg.Done()
}

func smoker(id int, signals, ingredients [3]chan int) {
	count := 0
	for range signals[id] {
		select {
		case <-ingredients[paper]:
		case <-ingredients[grass]:
		case <-ingredients[match]:
		}
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ingredients[paper]:
		case <-ingredients[grass]:
		case <-ingredients[match]:
		}
		time.Sleep(time.Millisecond * 500)
		count++
		fmt.Printf("%v%s smokes %v cigarettes\n",
			strings.Repeat("\t", 3+6*id),
			smokers[id],
			count)
	}
	wg.Done()
}

func main() {
	var ingredients [3]chan int
	var signals [3]chan int

	wg.Add(4)
	for i := range smokers {
		ingredients[i] = make(chan int)
		signals[i] = make(chan int)
	}

	for i := range smokers {
		go smoker(i, signals, ingredients)
	}
	go arbitrate(signals, ingredients)
	wg.Wait()
}
