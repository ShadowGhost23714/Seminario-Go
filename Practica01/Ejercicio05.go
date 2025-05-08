package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	switch {
	case (-99999 < i) && (i < -18):
		fmt.Println(i * -1)
	case (-18 <= i) && (i <= -1):
		fmt.Println(i % 4)
	case (1 <= i) && (i < 20):
		fmt.Println(i * i)
	case (20 <= i) && (i < 99999):
		fmt.Println(i * -1)
	}
}
