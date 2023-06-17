// https://leetcode.com/problems/plus-one/

package plus_one

func plusOne(digits []int) []int {
	reversedDigits := reverse(digits)
	toAdd := 1
	for k, v := range reversedDigits {
		nv := v
		if toAdd > 0 {
			nv += toAdd
			toAdd = 0
		}
		if nv > 9 {
			toAdd = nv - 9
			nv = 0
		}
		reversedDigits[k] = nv
	}
	if toAdd > 0 {
		reversedDigits = append(reversedDigits, []int{toAdd}...)
	}
	return reverse(reversedDigits)
}

func reverse(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
