package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadWords(fileName string) [][]string {
	ReadFile, _ := os.Open(fileName)
	FileScanner := bufio.NewScanner(ReadFile)
	var Line [][]string
	for FileScanner.Scan() {
		Line = append(Line, strings.Split(FileScanner.Text(), ""))
	}
	var AllLeters [][]string
	for i := 32; i < 126; i++ {
		AllLeters = append(AllLeters, ReadLetter(byte(i), "standard"))
	}
	for i := 32; i < 126; i++ {
		AllLeters = append(AllLeters, ReadLetter(byte(i), "shadow"))
	}
	for i := 32; i < 126; i++ {
		AllLeters = append(AllLeters, ReadLetter(byte(i), "thinkertoy"))
	}
	var words []string
	for len(Line[0]) > 0 {
		n := 0
		find := false
		for i := 0; i < len(AllLeters); i++ {
			if i == 94 || i == 188 {
				n++
			}
			if Hi(Line, AllLeters[i]) {
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
	for i := 0; i < len(words); i++ {
		fmt.Print(words[i])
	}
	ReadFile.Close()
	return Line
}

func Hi(Line [][]string, AllLeters []string) bool {
	for n := 0; n < len(AllLeters[0]) && n < len(Line[0]); n++ {
		for w := 0; w < len(Line); w++ {
			if Line[w][n] != string(AllLeters[w][n]) {
				return false
			}
		}
	}
	return true
}

func Terimming(startIndex, EndIndex int, Line [][]string) {
	for i := 0; i < len(Line); i++ {
		Line[i] = Line[i][startIndex:EndIndex]
	}
}
