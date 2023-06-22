// https://leetcode.com/problems/multiply-strings/

package multiply_strings

import "math/big"

func multiply(num1 string, num2 string) string {
	x := new(big.Int)
	y := new(big.Int)
	result := new(big.Int)

	x.SetString(num1, 10)
	y.SetString(num2, 10)
	result.Mul(x, y)

	return result.Text(10)
}
