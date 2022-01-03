package main

import (
	"fmt"
)

func main() {

	fmt.Println("Вывод ёлочки.")

	fmt.Println("Введите высоту ёлочки:")
	var height int
	fmt.Scan(&height)

	star := 1

	for aHeight := 0; aHeight < height; aHeight++ {

		for n := 0; n < height-aHeight; n++ {
			fmt.Print(" ")
		}

		for i := 0; i < star; i++ {
			fmt.Print("*")
		}

		star += 2
		fmt.Println("")
	}
}
