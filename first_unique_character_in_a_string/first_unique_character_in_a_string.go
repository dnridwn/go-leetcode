// https://leetcode.com/problems/first-unique-character-in-a-string/

package first_unique_character_in_a_string

func firstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		found := 0
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				found += 1
			}
		}
		if found == 1 {
			return i
		}
	}
	return -1
}
