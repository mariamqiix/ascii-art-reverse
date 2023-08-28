package ascii

import (
	"bufio"
	"os"
)

func ReadLetter(Text1 byte, fileName string) []string {
	//buffio object, to open and read the file
	ReadFile, _ := os.Open(fileName + ".txt")
	FileScanner := bufio.NewScanner(ReadFile)
	var Letter []string
	stop := 1
	i := 0
	a := (int(Text1)-32)*9 + 2
	for FileScanner.Scan() {
		i++
		if i >= a {
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
