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

func main() {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	go responseSize("https://www.info.unlp.edu.ar")
	fmt.Scanln()
}
