package iteration

import "strings"

func Repeat(s string, repeatCount int) string {
	var output strings.Builder
	for i := 0; i < repeatCount; i++ {
		output.WriteString(s)
	}
	return output.String()
}
