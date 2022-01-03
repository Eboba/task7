package main

import (
	"fmt"
)

func main() {

	fmt.Println("Зеркальные билеты")

	ticket := 100000
	luckyTicket := 0

	for ticket <= 999999 {

		var a, b, c, d, e, f = ticket / 100000 % 10, ticket / 10000 % 10, ticket / 1000 % 10, ticket / 100 % 10, ticket / 10 % 10, ticket % 10

		if a == f && b == e && c == d {
			luckyTicket++
			fmt.Println("Зеркальный билет:", ticket)
		}

		ticket++
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:")
	fmt.Println(luckyTicket)
}
