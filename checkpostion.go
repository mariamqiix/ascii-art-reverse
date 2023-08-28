package ascii

import (
	"os"
	"strings"
)

func CheckPostion(Text1, color, letterLenA, validation string) ([]int, []string) {
	found2 := false
	after2 := ""
	var nbr []int
	var colors []string
	TheWord := Text1
	TheWord2 := Text1
	for i := 0; i < len(Text1); i++ {
		_, after, found := strings.Cut(TheWord, letterLenA)
		if found == true {
			color = color
			nbr = append(nbr, (len(Text1) - len(after) - len(letterLenA)))
			colors = append(colors, color)
			TheWord = after
		} else if validation == "colorW2L" || validation == "colorW2LF" {
			_, after2, found2 = strings.Cut(TheWord2, os.Args[4])
			if found2 == true {
				color = CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[3], "--color=")))
				nbr = append(nbr, (len(Text1) - len(after2) - len(os.Args[4])))
				colors = append(colors, color)
			}
			TheWord2 = after2
		}
	}
	SortWordArr(nbr, colors)
	return nbr, colors
}

func SortWordArr(a []int, b []string) {
	temp2 := ""
	for j := 0; j < len(a)-1; j++ {
		temp := a[j]
		for i := j + 1; i < len(a); i++ {
			if a[i] < a[j] {
				temp = a[i]
				temp2 = b[i]
				a[i] = a[j]
				b[i] = b[j]
				a[j] = temp
				b[j] = temp2
			}
		}
	}
}
