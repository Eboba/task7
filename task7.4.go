package main

import (
	"fmt"
)

func main() {

	fmt.Println("Счастливые билеты.")

	ticket := 100000

	var (
		luckyTicket int
		nextLuck    int
	)

	for ticket <= 999999 {

		var a, b, c, d, e, f = ticket / 100000 % 10, ticket / 10000 % 10, ticket / 1000 % 10, ticket / 100 % 10, ticket / 10 % 10, ticket % 10

		if a+b+c == d+e+f {

			if nextLuck < luckyTicket {
				nextLuck = luckyTicket
			}

			luckyTicket = 0
		}

		luckyTicket++
		ticket++
	}

	fmt.Println("Минимальное количество билетов, которое нужно купить, чтобы среди них оказался счастливый:")
	fmt.Println(nextLuck)
}
