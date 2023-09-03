package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	FirstWord := true
	validation, Text, fileName, Letter, align, reverseFileName, color := ascii.Validation()
	WordsInArr := strings.Split(Text, "\\n")
	if strings.Contains(validation, "R") {
		ascii.PrintWithReverse(reverseFileName)
		os.Exit(0)
	} else if strings.Contains(validation, "output") {
		var test [][]string
		ascii.WriteFile(test, FirstWord, validation, align, 1)
	}

	if !(ascii.CheckTextSizeWithWidth(WordsInArr, strings.ToLower(fileName))) {
		fmt.Println("too much words , write less")
		os.Exit(0)
	} else if ascii.OnlyContains(Text, "\\n") {
		WordsInArr = WordsInArr[:len(WordsInArr)-1]
	}

	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		Text1 := strings.ReplaceAll(WordsInArr[l], "\\t", "   ")
		for j := 0; j < len(Text1); j++ {
			Words = append(Words, ascii.ReadLetter(Text1[j], strings.ToLower(fileName)))
		}
		if len(Text1) != 0 || strings.Contains(validation, "output") {
			count := strings.Count(Text1, " ")
			if strings.Contains(validation, "output") {
				ascii.WriteFile(Words, FirstWord, validation, align, count)
				FirstWord = false
			} else if strings.Contains(validation, "color") {
				ascii.PrintWithColor(Words, color, Text1, Letter, validation, align, count)
			}
		} else if !strings.Contains(validation, "output") {
			fmt.Println()
		}
	}
}
