package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSize(url string) {
	defer wg.Done()

	fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url, len(body))
}

var urls = []string{
	"https://www.golangprograms.com",
	"https://coderwall.com",
	"https://stackoverflow.com",
	"https://www.info.unlp.edu.ar",
}

func main() {
	for _, url := range urls {
		wg.Add(1)
		go responseSize(url)
	}
	wg.Wait()
}
