// https://leetcode.com/problems/palindrome-number/

package palindrome_number

import (
	"strconv"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	return xStr == reverseString(xStr)
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
