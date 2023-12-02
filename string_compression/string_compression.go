package string_compression

import (
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	mapCount := map[byte]int{}

	var builder strings.Builder

	for i, char := range chars {
		if i > 0 && char != chars[i-1] {
			mapCount[char] = 0
		}
		mapCount[char]++

		if i == len(chars)-1 || chars[i+1] != char {
			builder.WriteByte(char)
			if count := mapCount[char]; count > 1 {
				numStr := strconv.Itoa(count)
				builder.WriteString(numStr)
			}
		}
	}

	result := builder.String()
	copy(chars, []byte(result))

	return len(result)
}
