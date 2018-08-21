package solutions

func twoSum(nums []int, target int) []int {
	invertedIndex := make(map[int]int)
	for i, x := range nums {
		j, seen := invertedIndex[target-x]
		if seen {
			return []int{j, i}
		}
		invertedIndex[x] = i
	}
	return []int{}
}
