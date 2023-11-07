// https://leetcode.com/problems/daily-temperatures/

package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	var result []int
	for k, temp := range temperatures {
		result = append(result, 0)
		for i := k; i < len(temperatures); i++ {
			if temp < temperatures[i] {
				result[k] = i - k
				break
			}
		}
	}
	return result
}
