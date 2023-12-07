package string

import (
	"fmt"
	"strings"
)

func main() {
	input := "#please###show##me#code"
	output := extractWords(input)
	fmt.Println(output)
}

func extractWords(input string) string {
	split := strings.Split(input, "#")
	var words []string
	for _, word := range split {
		if word == "" {
			continue
		}
		words = append(words, word)
	}
	reverse(words)
	return strings.Join(words, " ")
}

func reverse(arr []string) {
	if len(arr) == 0 {
		return
	}
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
