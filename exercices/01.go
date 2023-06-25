package exercices

import "strconv"

func ConvertToInt(text string) (int, string) {
	i, _ := strconv.Atoi(text)
	if i > 100 {
		return i, "Greater than 100"
	}
	return i, "less than 100"
}
