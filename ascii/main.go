package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation, Text, fileName, Letter, align, reverseFileName, color, fileOutputName := ascii.Validation()
	WordsInArr := strings.Split(Text, "\\n")
	if strings.Contains(validation, "R") {
		ascii.PrintWithReverse(reverseFileName)
		os.Exit(0)
	} else if ascii.OnlyContains(Text, "\\n") {
		WordsInArr = WordsInArr[:len(WordsInArr)-1]
	}
	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		count := strings.Count(WordsInArr[l], " ")
		for j := 0; j < len(WordsInArr[l]); j++ {
			Words = append(Words, ascii.ReadLetter(WordsInArr[l][j], strings.ToLower(fileName)))
		}
		if strings.Contains(validation, "output") {
			ascii.WriteFile(Words, validation, align, count, fileOutputName)
		} else if strings.Contains(validation, "color") {
			if !(ascii.CheckTextSizeWithWidth(WordsInArr, strings.ToLower(fileName))) {
				fmt.Println("too much words, write less")
				os.Exit(0)
			} else if len(WordsInArr[l]) != 0 {
				ascii.PrintWithColor(Words, color, WordsInArr[l], Letter, validation, align, count)
			} else {
				fmt.Println()
			}
		}
	}
}
