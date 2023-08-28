package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() string {
	val := "yes"
	colorFlag := false
	outputFlag := false
	reverseFlag := false
	fsFlag := false
	alignFlag := false
	stringFlag := false
	for i := 1; i < len(os.Args); i++ {
		if strings.Index(os.Args[i], "--output=") == 0 && !stringFlag && !outputFlag {
			OutputValdation(i)
			val = "output"
			if colorFlag {
				val += "C"
			}
			outputFlag = true
		} else if strings.Index(os.Args[i], "--color=") == 0 && !stringFlag && !colorFlag {
			ColorValdation(i)
			if !outputFlag {
				val = "color"
			}
			colorFlag = true
		} else if strings.Index(os.Args[i], "--align=") == 0 && !stringFlag && !alignFlag {
			JustifyValidation(i)
			alignFlag = true
		} else if strings.Index(os.Args[i], "--reverse=") == 0 && !stringFlag && !reverseFlag {
			// JustifyValidation(i)
			reverseFlag = true
		} else if !stringFlag {
			CheckLetter(os.Args[i])
			stringFlag = true
		} else if stringFlag && !fsFlag {
			fsFlag = true
		} else {
			Error()
		}
	}
	if !stringFlag && !reverseFlag {
		Error()
	} else if (fsFlag || reverseFlag || alignFlag) && val != "yes" {
		val += "W"
		if fsFlag {
			val += "F"
		}
		if alignFlag {
			val += "J"
		}
	}
	if reverseFlag {
		val = "Reverse"
	}

	return val
}

func ColorValdation(i int) {
	if CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[i], "--color="))) == "NO" {
		fmt.Println("color does not exist")
		Error()
	}
}

func OutputValdation(i int) {
	if len(os.Args[i]) > 9 && strings.Index(os.Args[i], ".txt") != -1 && len(os.Args[i]) == strings.Index(os.Args[i], ".txt")+4 {
	} else {
		fmt.Println("The file name is incorrect")
		Error()
	}
}

func LetterExistnce(a, b int) {
	if strings.Index(os.Args[a], os.Args[b]) == -1 || len(os.Args[b]) == 0 {
		fmt.Println("letters to be colored does not exist in the word")
		Error()
	}
}

func JustifyValidation(i int) {
	align := strings.ToLower(strings.TrimPrefix(os.Args[i], "--align="))
	if !(align == "justify" || align == "left" || align == "right" || align == "center") {
		fmt.Println("align is incorrect, you can use (left, right, center, justify) only")
		Error()
	}
}

func CheckFont(s string) bool {
	FontType := strings.ToLower(s)
	if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
		return false
	}
	return true
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("Example: go run . --align=right something standard")
	os.Exit(0)
}

func CheckLetter(s string) {
	for g := 0; g < len(s); g++ {
		if s[g] > 126 || s[g] < 32 {
			fmt.Println("ERROR: ascii letters only")
			os.Exit(0)
		}
	}
}
