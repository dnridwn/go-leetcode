package move_zeroes

func moveZeroes(nums []int) {
	nNums := &nums
	kSearch := 0
	countScreening := 0
	for countScreening < 2 {
		n := (*nNums)[kSearch]
		if n == 0 {
			for i := kSearch; i < len((*nNums))-1; i++ {
				(*nNums)[i] = (*nNums)[i+1]
			}
			(*nNums)[len((*nNums))-1] = n
		}
		kSearch += 1
		if kSearch >= len((*nNums))-1 {
			kSearch = 0
			countScreening += 1
		}
	}
}
