// https://leetcode.com/problems/fizz-buzz

package fizz_buzz

import "strconv"

func fizzBuzz(n int) (result []string) {
	for i := 1; i <= n; i++ {
		answer := strconv.Itoa(i)
		if (i%3 == 0) && (i%5 == 0) {
			answer = "FizzBuzz"
		} else if i%3 == 0 {
			answer = "Fizz"
		} else if i%5 == 0 {
			answer = "Buzz"
		}
		result = append(result, answer)
	}

	return
}
