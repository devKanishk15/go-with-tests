package iterations

import "strings"

// const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}


func ExampleRepeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}