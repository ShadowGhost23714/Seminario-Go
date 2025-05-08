package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string) {
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
		go responseSize(url)
	}
	fmt.Scanln()
}
