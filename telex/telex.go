package telex

import "strings"

func Transform(input string) string {
	var result strings.Builder

	arr := strings.Split(input, " ")
	l := len(arr)
	for i, item := range arr {
		result.WriteString(transformWord(item))
		if i < l-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}
