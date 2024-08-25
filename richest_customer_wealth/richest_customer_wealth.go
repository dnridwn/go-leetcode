package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	highest := 0
	for _, account := range accounts {
		sum := sum(account)
		if sum > highest {
			highest = sum
		}
	}

	return highest
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
