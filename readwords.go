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

	width := TakingTheWidth(Line)

	var AllLeters [][]string
	for i := 32; i < 126; i++ {
		AllLeters = append(AllLeters, ReadLetter(byte(i), "standard"))
	}
	var num []int
	end := true
	for end {
		for i := 0; i < 94; i++ {
			if Hi(width, Line, AllLeters[i]) {
				num = append(num, i+32)
				fmt.Println(i + 32)
				var test [][]string
				test = append(test, AllLeters[i])
				startIndex := TakingTheWidth(test) + 1
				EndIndex := TakingTheWidth(Line)
				if EndIndex-startIndex > 1 {
					Line = Terimming(startIndex, EndIndex, Line)
				} else {
					end = false
				}
				break
			}
		}
	}
	fmt.Print(num)
	fmt.Println(string(num[0]))
	ReadFile.Close()
	return Line
}

func Hi(width int, Line [][]string, AllLeters []string) bool {
	var test [][]string
	test = append(test, AllLeters)

	for n := 0; n < TakingTheWidth(test) && n < TakingTheWidth(Line); n++ {
		for w := 0; w < len(Line); w++ {
			if Line[w][n] != string(AllLeters[w][n]) {
				return false
			}
		}
	}
	return true
}

func Terimming(startIndex, EndIndex int, Line [][]string) [][]string {
	var Line2 [][]string
	for i := 0; i < len(Line); i++ {
		Line2 = append(Line2, Line[i][startIndex:EndIndex])
	}
	return Line2
}

func TakingTheWidth(Line [][]string) int {
	width := 0
	for _, row := range Line {
		if len(row) > width {
			width = len(row)
		}
	}
	return width
}
