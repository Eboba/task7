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

	cellWt := " "
	cellBl := "*"

	for aHeight < height {

		for aWidth < width {
			fmt.Print(cellWt)
			aWidth++

			if aWidth == width {
				break
			}

			fmt.Print(cellBl)
			aWidth++
		}

		cellWt, cellBl = cellBl, cellWt // Вот и Ваша подсказка помогла из 3 модуля, сделал до того как видео посмотрел =)
		fmt.Println()
		aHeight++
		aWidth = 0
	}
}
