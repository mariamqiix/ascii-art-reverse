package ascii

import (
	"fmt"
	"strings"
)

func PrintWithJustify(Words [][]string, align string, n, w, count int) string {
	textSize, spaceSize, width := "", 1, width()-1
	for n := 0; n < len(Words); n++ {
		textSize += Words[n][3]
	}
	if len(textSize) < width {
		spaceSize = width - len(textSize)
	}
	if align == "right" && n == 0 {
		return fmt.Sprintf(strings.Repeat(" ", spaceSize) + Words[n][w])
	} else if align == "center" && n == 0 {
		return fmt.Sprintf(strings.Repeat(" ", spaceSize/2) + Words[n][w])
	} else if align == "justify" && n != 0 && Words[n][3] == "      " && Words[n][5] == "      " && Words[n][7] == "      " {
		if count == 0 {
			count = 1
		}
		return fmt.Sprintf(strings.Repeat(" ", spaceSize/count) + Words[n][w])
	}
	return fmt.Sprintf(Words[n][w])
}
