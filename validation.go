package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() (string, string, string, string, string, string, string) {
	val := "color"
	colorFlag, outputFlag, reverseFlag, fsFlag, alignFlag, stringFlag, colorWLflag, colorW2Lflag := false, false, false, false, false, false, false, false
	TheText, Letter, reverseFileName, fileName, align, color := "", "", "", "standard", "left", "\033[0m"
	for i := 1; i < len(os.Args); i++ {
		if (len(os.Args) == 6 || len(os.Args) == 7) && strings.Index(os.Args[1], "--color=") == 0 && strings.Index(os.Args[3], "--color=") == 0 && i == 1 {
			val = "colorW2L"
			colorW2Lflag = true
			if strings.ToLower(os.Args[3]) == strings.ToLower(os.Args[1]) { //check if u cant print with same color
				fmt.Println("The colors should be different")
				Error()
			}
			ColorValdation(3)
			ColorValdation(1)
			LetterExistnce(os.Args[5], os.Args[4])
			i += 3
			color = ColorValdation(1)
			Letter = os.Args[2]
		} else if strings.Index(os.Args[i], "--output=") == 0 && !stringFlag && !outputFlag {
			val, outputFlag = "output", true
			if (strings.Index(os.Args[i], ".txt") == -1) || len(os.Args[i]) != strings.Index(os.Args[i], ".txt")+4 {
				fmt.Println("The file name is incorrect")
				Error()
			}
		} else if strings.Index(os.Args[i], "--color=") == 0 && !stringFlag && !colorFlag && !colorW2Lflag {
			val, colorFlag = "color", true
			color = ColorValdation(i)
		} else if strings.Index(os.Args[i], "--align=") == 0 && !stringFlag && !alignFlag {
			align, alignFlag = strings.ToLower(strings.TrimPrefix(os.Args[i], "--align=")), true
			if !(align == "justify" || align == "left" || align == "right" || align == "center") {
				fmt.Println("align is incorrect, you can use (left, right, center, justify) only")
				Error()
			}
		} else if strings.Index(os.Args[i], "--reverse=") == 0 && !stringFlag && !reverseFlag {
			val, reverseFlag = "Reverse", true
			if len(os.Args[i]) <= 10 {
				fmt.Println("The file name is missing")
				Error()
			}
			reverseFileName = os.Args[i][10:]
		} else if colorFlag && i+1 < len(os.Args) && !colorWLflag && os.Args[i+1] != "standard" && os.Args[i+1] != "shadow" && os.Args[i+1] != "thinkertoy" {
			colorWLflag = true
			Letter = os.Args[i]
		} else if !stringFlag && !reverseFlag {
			TheText, stringFlag = os.Args[i], true
			for g := 0; g < len(TheText); g++ {
				if TheText[g] > 126 || TheText[g] < 32 {
					fmt.Println("ERROR: ascii letters only")
					os.Exit(0)
				}
			}
		} else if stringFlag && !fsFlag {
			fileName, fsFlag = strings.ToLower(os.Args[i]), true
			if fileName != "standard" && fileName != "shadow" && fileName != "thinkertoy" {
				Error()
			}
		} else {
			Error()
		}
	}
	if Letter != "" {
		LetterExistnce(TheText, Letter)

	}
	if !stringFlag && !reverseFlag {
		Error()
	} else if outputFlag && colorFlag {
		val = "output"
		fmt.Println("The colored text can't be print inside the file")
	}
	return val, TheText, fileName, Letter, align, reverseFileName, color
}

func ColorValdation(i int) string {
	color := CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[i], "--color=")))
	if color == "NO" {
		fmt.Println("color does not exist")
		Error()
	}
	return color
}

func LetterExistnce(a, b string) {
	if strings.Index(a, b) == -1 || len(b) == 0 {
		fmt.Println("letters to be colored does not exist in the word")
		Error()
	}
}

func Error() {
	fmt.Println("Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>")
	os.Exit(0)
}
