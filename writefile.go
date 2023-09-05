package ascii

import (
	"fmt"
	"os"
)

func WriteFile(s [][]string, validation, align string, count int, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error \n", err)
	} else {
		if len(s) != 0 {
			for w := 0; w < 8; w++ {
				for n := 0; n < len(s); n++ {
					file.WriteString(PrintWithJustify(s, align, n, w, count))
				}
				if w+1 != 8 {
					file.WriteString("\n")
				}
			}
			file.WriteString("\n")
		} else {
			file.WriteString("\n")
		}
	}
	file.Close()
}
