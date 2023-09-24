package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() (string, string, string, string, string, string, string, string) {
	val := "color"
	colorFlag, outputFlag, reverseFlag, fsFlag, alignFlag, stringFlag, colorWLflag, colorW2Lflag := false, false, false, false, false, false, false, false
	TheText, Letter, reverseFileName, fileName, align, color, fileOutputName := "", "", "", "standard", "left", "\033[0m", ""
	for i := 1; i < len(os.Args); i++ {
		if (len(os.Args) == 6 || len(os.Args) == 7) && strings.Index(os.Args[1], "--color=") == 0 && strings.Index(os.Args[3], "--color=") == 0 && i == 1 {
			val, colorW2Lflag = "colorW2L", true
			LetterExistnce(os.Args[5], os.Args[4])
			i += 3
			color, Letter = ColorValdation(1), os.Args[2]
		} else if strings.Index(os.Args[i], "--output=") == 0 && !stringFlag && !outputFlag {
			val, outputFlag, fileOutputName = "output", true, os.Args[i][9:]
			os.Remove(fileOutputName)
			if (strings.Index(os.Args[i], ".txt") == -1) || len(os.Args[i]) != strings.Index(os.Args[i], ".txt")+4 {
				fmt.Println("The file name is incorrect")
				Error()
			}
		} else if strings.Index(os.Args[i], "--color=") == 0 && !stringFlag && !colorFlag && !colorW2Lflag {
			colorFlag, color = true, ColorValdation(i)
		} else if strings.Index(os.Args[i], "--align=") == 0 && !stringFlag && !alignFlag {
			align, alignFlag = strings.ToLower(os.Args[i][8:]), true
			if !(align == "justify" || align == "left" || align == "right" || align == "center") {
				fmt.Println("align is incorrect, you can use (left, right, center, justify) only")
				Error()
			}
		} else if strings.Index(os.Args[i], "--reverse=") == 0 && !stringFlag && !reverseFlag  {
			val, reverseFlag, reverseFileName = "Reverse", true, os.Args[i][10:]
			if (strings.Index(os.Args[i], ".txt") == -1) || len(os.Args[i]) != strings.Index(os.Args[i], ".txt")+4 {
				fmt.Println("The file name is incorrect")
				Error()
			} else if len(os.Args) != 2 {
				Error()
			} 
		} else if colorFlag && i+1 < len(os.Args) && !colorWLflag && os.Args[i+1] != "standard" && os.Args[i+1] != "shadow" && os.Args[i+1] != "thinkertoy" {
			colorWLflag, Letter = true, os.Args[i]
		} else if !stringFlag && !reverseFlag {
			TheText, stringFlag = strings.ReplaceAll(os.Args[i], "\\t", "   "), true
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
	if alignFlag {
		TheText = strings.Join(strings.Fields(TheText)," ")
	}
	if colorW2Lflag || colorWLflag {
		LetterExistnce(TheText, Letter)
	}
	if !stringFlag && !reverseFlag {
		Error()
	} else if outputFlag && colorFlag {
		fmt.Println("The colored text can't be print inside the file")
	}
	return val, TheText, fileName, Letter, align, reverseFileName, color, fileOutputName
}

func ColorValdation(i int) string {
	color := CheckColor(os.Args[i])
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
