package arrays

func GetWorkHours(schedule []int) []int {

	if len(schedule) < 5 {
		return []int{}
	}

	switch len(schedule) {
	case 7:
		return schedule[:len(schedule)-2]
	case 6:
		return schedule[:len(schedule)-1]
	case 5:
		return schedule[:]
	}

	return nil
}
