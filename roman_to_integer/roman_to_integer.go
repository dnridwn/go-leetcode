// https://leetcode.com/problems/roman-to-integer/

package roman_to_integer

func romanToInt(s string) int {
	var (
		result       int            = 0
		romanSymbols map[string]int = map[string]int{
			"I":  1,
			"IV": 4,
			"V":  5,
			"IX": 9,
			"X":  10,
			"XL": 40,
			"L":  50,
			"XC": 90,
			"C":  100,
			"CD": 400,
			"D":  500,
			"CM": 900,
			"M":  1000,
		}
	)

	for i := 0; i < len(s); i++ {
		symbol := string(s[i])
		if (symbol == "I" || symbol == "X" || symbol == "C") && i < len(s)-1 {
			doubleSymbol := string(s[i : i+2])
			if v, ok := romanSymbols[doubleSymbol]; ok {
				result += v
				i++
				continue
			}
		}

		result += romanSymbols[symbol]
	}

	return result
}
