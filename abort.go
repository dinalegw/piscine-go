package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	// simple bubble sort for 5 elements
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// median of 5 numbers is the middle one (index 2)
	return nums[2]
}
