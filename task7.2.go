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
		aHeight int
		aWidth  int
	)

	for aHeight < height {

		for aWidth < width {
			fmt.Print(" ")
			aWidth++

			if aWidth == width {
				break
			}

			fmt.Print("*")
			aWidth++

		}

		fmt.Println()
		aHeight++
		aWidth = 0

		for aWidth < width {
			fmt.Print("*")
			aWidth++

			if aWidth == width {
				break
			}

			fmt.Print(" ")
			aWidth++

		}

		fmt.Println()
		aHeight++
		aWidth = 0
	}
}
