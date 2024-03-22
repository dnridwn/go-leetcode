package can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 {
		return false
	}

	if len(flowerbed) == 1 {
		return flowerbed[0] == 0
	}

	if len(flowerbed) == 2 {
		return flowerbed[0] == 0 && flowerbed[1] == 0 && n <= 1
	}

	x := 0
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				x++
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				x++
			}
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				x++
			}
		}
	}

	return x >= n
}
