package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation := ascii.Validation()
	index := 1
	if !strings.Contains(validation, "W") && validation != "yes" {
		index++
	} else if !strings.Contains(validation, "F") && validation != "yes" {
		index = len(os.Args) - 1
	} else if strings.Contains(validation, "F") {
		index = len(os.Args) - 2
	}

	if strings.Contains(validation, "R") {
		reverseWord := strings.ToLower(ascii.ReturnFlag("--reverse="))
		ascii.ReadWords(reverseWord)
		return

	}
	WordsInArr := strings.Split(os.Args[index], "\\n")
	fileName := "standard"
	if (len(os.Args) == 3 && validation == "yes") || (len(os.Args) == 4 && !strings.Contains(validation, "W")) || (strings.Contains(validation, "F")) {
		fileName = strings.ToLower(os.Args[index+1])
		if !ascii.CheckFont(fileName) {
			ascii.Error()
		}
	}

	if !(ascii.CheckTextSizeWithWidth(WordsInArr, fileName)) {
		fmt.Println("too much words , write less")
		return
	}

	if ascii.OnlyContains(os.Args[index], "\\n") {
		WordsInArr = WordsInArr[:len(WordsInArr)-1]
	}

	FirstWord := true
	align := "left"
	if strings.Contains(validation, "J") {
		align = strings.ToLower(ascii.ReturnFlag("--align="))
	}

	if strings.Contains(validation, "output") {
		var test [][]string
		ascii.WriteFile(test, FirstWord, validation, align, 1)
	}
	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		Text1 := strings.ReplaceAll(WordsInArr[l], "\\t", "   ")
		for j := 0; j < len(Text1); j++ {
			Words = append(Words, ascii.ReadLetter(Text1[j], fileName))
		}
		if len(Text1) != 0 || strings.Contains(validation, "output") {
			count := strings.Count(Text1, " ")
			if strings.Contains(validation, "output") {
				if l == 0 && strings.Contains(validation, "C") {
					fmt.Println("The colored text can't be print inside the file")
				}
				ascii.WriteFile(Words, FirstWord, validation, align, count)
				FirstWord = false
			} else if strings.Contains(validation, "color") {
				letter1 := "NO!!-"
				color := ascii.CheckColor(strings.ToLower(ascii.ReturnFlag("--color=")))
				if validation == "colorW2L" || validation == "colorW2LF" || validation == "colorWLJ" || validation == "colorWLJF" {
					letter1 = os.Args[2]
				} else if strings.Contains(validation, "L") {
					letter1 = os.Args[index-1]
				}
				ascii.PrintWithColor(Words, color, Text1, letter1, validation, align, count)
			} else {
				for w := 0; w < 8; w++ {
					for n := 0; n < len(Words); n++ {
						fmt.Print(ascii.PrintWithJustify(Words, align, n, w, count))
					}
					if w+1 != 8 {
						fmt.Println()
					}
				}
				fmt.Println()
			}
		} else if !strings.Contains(validation, "output") {
			fmt.Println()
		}
	}
}
