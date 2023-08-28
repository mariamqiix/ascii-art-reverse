package ascii

func OnlyContains(s, sep string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "\\" {
			if string(s[i+1]) == "n" {
				if i != len(s)-3 {
					i++
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
