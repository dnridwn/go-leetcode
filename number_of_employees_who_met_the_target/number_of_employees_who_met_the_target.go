package number_of_employees_who_met_the_target

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	r := 0
	for _, n := range hours {
		if n >= target {
			r++
		}
	}

	return r
}
