package solutions

func firstMissingPositive(nums []int) int {
	// do a scan recording seen numbers in place. 0 means haven't seen it, 1 means seen
	n := len(nums)
	for i := 0; i < n; i++ {
		// this is 1 + the index we want to mark
		x := nums[i]
		switch {
		case x < 1 || x > n:
			nums[i] = 0
		case x-1 == i:
			nums[i] = 1
		case x-1 < i:
			nums[x-1] = 1
			nums[i] = 0
		case x == nums[x-1]:
			nums[i] = 0
		default:
			nums[i] = nums[x-1]
			nums[x-1] = x
			i--
		}
	}

	// either find the first 1-based array index we haven't seen or past length
	for i, f := range nums {
		if f == 0 {
			return i + 1
		}
	}
	return n + 1
}
