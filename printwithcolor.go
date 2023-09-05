package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintWithColor(Words [][]string, color1, Text, letter1, validation, align string, count int) {
	positions := CheckColors(Text, color1, letter1, validation)
	for w := 0; w < 8; w++ {
		for n := 0; n < len(Words); n++ {
			fmt.Print(positions[n], PrintWithJustify(Words, align, n, w, count))
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func CheckColors(Text1, color, letter1, validation string) map[int]string {
	indexWcolor := make(map[int]string)
	for i := 0; i < len(Text1); i++ {
		if i+len(letter1) <= len(Text1) && Text1[i:i+len(letter1)] == letter1 && letter1 != "" {
			for j := 0; j < len(letter1); j++ {
				indexWcolor[i] = color
			}
			i += len(letter1) - 1
		} else if strings.Contains(validation, "2") && i+len(letter1) <= len(Text1) && Text1[i:i+len(os.Args[4])] == os.Args[4] {
			for j := 0; j < len(os.Args[4]); j++ {
				indexWcolor[i] = CheckColor(os.Args[3])
			}
			i += len(os.Args[4]) - 1
		} else if letter1 == "" {
			indexWcolor[i] = color
		} else {
			indexWcolor[i] = "\033[0m"
		}
	}
	return indexWcolor
}
