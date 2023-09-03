package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintWithColor(Words [][]string, color, Text1, letter1, validation, align string, count int) {
	colorB := color
	positions, colors := CheckPostion(Text1, colorB, letter1, validation)
	FlagB := false
	for w := 0; w < 8; w++ {
		index := 0
		stop := 0
		for n := 0; n < len(Words); n++ {
			if letter1 != "" {
				if len(positions) != 0 {
					if positions[index] == n {
						FlagB = true
					}
					if FlagB {
						colorB = colors[index]
						stop++
						if (stop == len(letter1) && colorB == color) || (strings.Contains(validation, "2") && stop == len(os.Args[4]) && color != colors[index]) {
							FlagB = false
							stop = 0
							if index+1 < len(positions) {
								index++
							}
						}
					} else {
						colorB = "\033[0m"
					}
				} else {
					colorB = "\033[0m"
				}
			}
			fmt.Print(colorB, PrintWithJustify(Words, align, n, w, count))
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}
