package ascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CheckTextSizeWithWidth(word []string) bool {
	width := width()
	for _, char := range word {
		wordWfont := ""
		for _, char2 := range char {
			wordWfont += ReadLetter(byte(char2))[0]
		}
		if len(wordWfont) > width {
			return false
		}
	}
	return true
}

func width() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(strings.Split(strings.TrimSpace(string(out)), " ")[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}
