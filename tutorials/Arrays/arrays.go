package arrays

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i < 0 {
		return nums
	}

	if i < len(nums) {
		nums[i] = val
	} else {
		return nums
	}

	return nums
}
