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
		opTicket    int
		tpTicket    int
	)

	for ticket <= 999999 {

		n := ticket

		for i := 0; i < 6; i++ {

			if i < 3 {
				tpTicket += n % 10
			} else {
				opTicket += n % 10
			}

			n /= 10
		}

		if opTicket == tpTicket {

			if nextLuck < luckyTicket {
				nextLuck = luckyTicket
			}

			luckyTicket = 0
		}

		opTicket, tpTicket = 0, 0

		luckyTicket++
		ticket++
	}

	fmt.Println("Минимальное количество билетов, которое нужно купить, чтобы среди них оказался счастливый:")
	fmt.Println(nextLuck)
}
