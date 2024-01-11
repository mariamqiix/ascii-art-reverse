package ascii

import (
	"fmt"
	"os"
)

func WriteFile(TheOutPutText [][]string, count int) {
	file, err := os.OpenFile(fileOutputName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)

	} else {
		
		for w := 0; w < 8 && len(TheOutPutText) != 0; w++ {

			for n := 0; n < len(TheOutPutText); n++ {
				file.WriteString(PrintWithJustify(TheOutPutText, n, w, count))
			}

			if w+1 != 8 {
				file.WriteString("\n")
			}
		}

		file.WriteString("\n")
	}

	file.Close()
}
