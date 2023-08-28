package ascii

import (
	"os"
	"strings"
)

func ReturnFlag(flag string) string {
	flagCut := "left"
	for x := 0; x < len(os.Args); x++ {
		if strings.Contains(os.Args[x], flag) {
			flagCut = strings.TrimPrefix(os.Args[x], flag)
			break
		}
	}
	return flagCut
}
