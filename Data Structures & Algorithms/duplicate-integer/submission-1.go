func hasDuplicate(nums []int) bool {
    mpp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mpp[nums[i]] += 1
	}

	for _, value := range(mpp) {
		if value >= 2 {
			return true
		}
	}

	return false
}
