package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLetter(Text1 byte) []string {
	var Letter []string
	ReadFile, err := os.Open(fileName + ".txt")

	if err != nil {
		fmt.Println(err)
		Error()
	}

	FileScanner := bufio.NewScanner(ReadFile)
	stop := 1
	letterLength := (int(Text1)-32)*9 + 2

	for i := 1; FileScanner.Scan(); i++ {
		if i >= letterLength {
			stop++
			Letter = append(Letter, FileScanner.Text())
			if stop > 8 {
				break
			}
		}
	}

	ReadFile.Close()
	return Letter
}
