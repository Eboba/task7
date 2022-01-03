package main

import (
	"fmt"
)

func main() {

	fmt.Println("Шахматная доска")

	fmt.Println("Введите ширину:")
	var width int
	fmt.Scan(&width)

	fmt.Println("Введите высоту:")
	var height int
	fmt.Scan(&height)

	var (
		aWidth  int
		aHeight int
	)

	for aHeight < height {

		fmt.Print(" ")
		aWidth++

		if aWidth == width {
			fmt.Println("")
			aWidth = 0
			aHeight++
		}

		fmt.Print("*")
		aWidth++

		if aWidth == width {
			fmt.Println("")
			aWidth = 0
			aHeight++
		}

	}
}
