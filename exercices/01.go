package exercices

import "strconv"

func ConvertToInt(text string) (int, string) {
	i, err := strconv.Atoi(text)
	if err != nil {
		return 0, "error ocurred: " + err.Error()
	}
	if i > 100 {
		return i, "Greater than 100"
	}
	return i, "less than 100"
}
