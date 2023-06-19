// https://leetcode.com/problems/add-digits/

package add_digits

import (
	"strconv"
	"strings"
)

func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	numInStr := strconv.Itoa(num)
	nums := strings.Split(numInStr, "")
	totalOfSum := 0
	for _, n := range nums {
		newN, _ := strconv.Atoi(n)
		totalOfSum += newN
	}
	return addDigits(totalOfSum)
}
