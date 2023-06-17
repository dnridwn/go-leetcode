// https://leetcode.com/problems/add-binary/

package add_binary

import "math/big"

func addBinary(a string, b string) string {
	num1 := new(big.Int)
	num2 := new(big.Int)
	num1.SetString(a, 2)
	num2.SetString(b, 2)

	result := new(big.Int)
	result.Add(num1, num2)

	return result.Text(2)
}
