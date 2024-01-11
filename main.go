package main

import (
	"ascii/func"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation, WordsInArr := ascii.Validation()

	if strings.Contains(validation, "R") {
		ascii.PrintWithReverse()
		os.Exit(0)

	}

	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		count := strings.Count(WordsInArr[l], " ")

		for j := 0; j < len(WordsInArr[l]); j++ {
			Words = append(Words, ascii.ReadLetter(WordsInArr[l][j]))
		}

		if strings.Contains(validation, "output") {
			ascii.WriteFile(Words, count)

		} else if strings.Contains(validation, "color") {

			if !(ascii.CheckTextSizeWithWidth(WordsInArr)) {
				fmt.Println("too much words, write less")
				os.Exit(0)

			} else if len(WordsInArr[l]) != 0 {
				ascii.PrintWithColor(Words, WordsInArr[l], count)

			} else {
				fmt.Println()
			}
		}
	}
}
