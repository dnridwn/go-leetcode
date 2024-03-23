package find_the_highest_altitude

func largestAltitude(gain []int) int {
	highest := 0
	alt := 0

	for _, v := range gain {
		alt = alt + v

		if alt > highest {
			highest = alt
		}
	}

	return highest
}
