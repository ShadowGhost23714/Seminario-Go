/*
4. ¿Cómo podría hacer para garantizar que el siguiente programa
imprima?

PING
PONG
PING
PONG
PING
PONG

package main

import "fmt"

func pxng(m chan string, str string) {
	m <- str
}

func main() {
	messages := make(chan string)
	for i := 0; i < 3; i++ {
		go pxng(messages, "PING")
		go pxng(messages, "PONG")
	}
	for i := 0; i < 6; i++ {
		fmt.Println(<-messages)
	}
}
*/

package main

import (
	"fmt"
)

// func pxng(m chan string, str string) {
// 	m <- str
// }
// func main() {
// 	messages := make(chan string)
// 	for i := 0; i < 10; i++ {
// 		go pxng(messages, "PING")
// 		fmt.Println(<-messages)
// 		go pxng(messages, "PONG")
// 		fmt.Println(<-messages)
// 	}
// }

func pxng(m chan string, str string) {
	m <- str
}
func main() {

	messagesPing := make(chan string)
	messagesPong := make(chan string)
	for i := 0; i < 5; i++ {
		go pxng(messagesPing, "PING")
		go pxng(messagesPong, "PONG")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-messagesPing)
		fmt.Println(<-messagesPong)
	}
}

// func pxng(m chan string, str string) {
// 	m <- str
// }
// func main() {
// 	messages := make(chan string)
// 	for i := 0; i < 5; i++ {
// 		go pxng(messages, "PING")
// 		go pxng(messages, "PONG")
// 	}
// 	for i := 0; i < 10; i++ {
// 		select {
// 			case ping
// 			case pong
// 		fmt.Println(<-messages)
// 	}
// }
