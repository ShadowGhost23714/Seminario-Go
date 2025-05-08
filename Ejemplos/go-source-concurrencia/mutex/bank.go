package main

import "fmt"

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	for i := 0; i < 10; i++ {
		go Deposit(100)
	}
	fmt.Println(Balance())
}
