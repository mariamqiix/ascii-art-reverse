package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintWithReverse() {

	var fileLine [][]string
	var AllLeters [][]string
	var words []string

	ReadFile, err := os.Open(reverseFileName)

	if err != nil {
		fmt.Println(err)
		Error()
	}

	FileScanner := bufio.NewScanner(ReadFile)

	

	for FileScanner.Scan() {
		fileLine = append(fileLine, strings.Split(FileScanner.Text(), ""))
	}

	
	for i := 32; i < 126; i++ {
		reverseFileName = "standard"
		AllLeters = append(AllLeters, ReadLetter(byte(i)))
	}

	for i := 32; i < 126; i++ {
		reverseFileName = "shadow"
		AllLeters = append(AllLeters, ReadLetter(byte(i)))
	}

	for i := 32; i < 126; i++ {
		reverseFileName = "thinkertoy"
		AllLeters = append(AllLeters, ReadLetter(byte(i)))
	}

	if len(fileLine)%8 != 0 {
		Error()
	}

	
	for i := 0; i < len(fileLine); i += 8 {
		var Line [][]string

		for j := i; j < i+8 && j < len(fileLine); j++ {
			Line = append(Line, fileLine[j])
		}

		for len(Line[0]) > 0 {
			n := 0
			find := false
			for i := 0; i < len(AllLeters); i++ {
				if i == 94 || i == 188 {
					n++
				}
				if CheckTheLetter(Line, AllLeters[i]) {
					words = append(words, string(i+32-(94*n)))
					Terimming(len(AllLeters[i][0]), len(Line[0]), Line)
					find = true
					break
				}
			}
			if !find {
				Error()
			}
		}

		words = append(words, "\n")
	}

	for i := 0; i < len(words)-1; i++ {
		fmt.Print(words[i])
	}

	fmt.Println()
	ReadFile.Close()

}

func CheckTheLetter(Line [][]string, AllLeters []string) bool {
	for n := 0; n < len(AllLeters[0]); n++ {

		for w := 0; w < len(Line); w++ {

			if len(Line[w]) == 0 || len(AllLeters[0]) > len(Line[0]) || Line[w][n] != string(AllLeters[w][n]) {
				return false
			}
		}
	}

	return true
}

func Terimming(startIndex, endIndex int, Line [][]string) {

	for i := 0; i < len(Line); i++ {
		if len(Line[i]) >= endIndex {
			Line[i] = Line[i][startIndex:endIndex]

		} else {
			Error()
		}
	}
}
