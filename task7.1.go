package main

import (
	"fmt"
)

func main() {

	fmt.Println("Зеркальные билеты")

	ticket := 100000
	luckyTicket := 0
	partTicket := 0

	for ticket <= 999999 {

		n := ticket

		for i := 0; i < 3; i++ {
			remainder := n % 10
			partTicket *= 10
			partTicket += remainder
			n /= 10
		}

		if ticket/1000%1000 == partTicket {
			luckyTicket++
			fmt.Println("Зеркальный билет:", ticket)
		}

		partTicket = 0
		ticket++
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:")
	fmt.Println(luckyTicket)
}
