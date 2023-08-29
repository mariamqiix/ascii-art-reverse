package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckColor(userValue string) string {
	Colors := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m",
		"blue": "\033[34m", "purple": "\033[35m", "cyan": "\033[36m", "white": "\033[37m",
		"black": "\033[30m", "orange": "\033[38;5;208m", "\033[0m":"\033[0m"}

	for color, value := range Colors {
		if color == userValue {
			return value
		} else if strings.Index(userValue, "rgb") == 0 && userValue[len(userValue)-1] == ')' {
			userValue := strings.ReplaceAll(userValue, " ", "")
			c := (strings.Split(userValue[4:len(userValue)-1], ","))
			for i := 0; i < len(c); i++ {
				check, err := strconv.Atoi(c[i])
				if err != nil || len(c) != 3 || check < 0 || check > 255 {
					return "NO"
				}
			}
			red, _ := strconv.Atoi(c[0])
			green, _ := strconv.Atoi(c[1])
			blue, _ := strconv.Atoi(c[2])
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
		} else if strings.Index(userValue, "#") == 0 && len(userValue) == 7 {
			for i := 1; i <= 6; i++ {
				if (userValue[i] >= '0' && userValue[i] <= '9') || (userValue[i] >= 'a' && userValue[i] <= 'f') {
				} else {
					return "NO"
				}
			}
			rgb, _ := strconv.ParseUint(userValue[1:], 16, 32)
			userValue = fmt.Sprintf("\033[38;2;%d;%d;%dm", int(rgb>>16&0xFF), int(rgb>>8&0xFF), int(rgb&0xFF))
			return userValue
		}
	}
	return "NO"
}
