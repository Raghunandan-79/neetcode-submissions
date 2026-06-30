func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n, m := len(s), len(t)
	mpp := make(map[byte]int)

	for i := 0; i < n; i++ {
		mpp[s[i]] += 1
	}

	for i := 0; i < m; i++ {
		mpp[t[i]] -= 1
	}

	for _, value := range(mpp) {
		if value != 0 {
		  return false
		}
	}

	return true
}