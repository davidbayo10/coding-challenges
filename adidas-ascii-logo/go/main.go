package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const COLUMNS = 3

func getSquare(width int) int {
	value := math.Sqrt(float64(width))

	return int(math.Round(value))
}

func createAdidasAsciiLogo(width int) string {
	if width < 2 {
		panic(errors.New("Error, minimun width is 2"))
	}

	var logo string
	square := getSquare(width)

	for i := 1; i <= COLUMNS; i++ {
		block := Block{COLUMNS, i, width}

		for j := 0; j < square; j++ {
			logo = fmt.Sprintf(
				"%s%s%s\n",
				logo,
				strings.Repeat(" ", j),
				block.Draw(square, "@"),
			)
		}
	}

	return logo
}

func main() {
	widths := []int{2, 3, 5, 7, 9, 16, 21}

	for _, width := range widths {
		fmt.Println(fmt.Sprintf("\nadidas (width %d)", width))
		fmt.Println("\n-------------------------------------------------------------")
		fmt.Println(fmt.Sprintf("\n%s\n\n", createAdidasAsciiLogo(width)))
	}
}
