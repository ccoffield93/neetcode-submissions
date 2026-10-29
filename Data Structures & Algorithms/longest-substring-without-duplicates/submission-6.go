func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	currMap := make(map[byte]int)
	l := 0
	r := 0
	longest := 0

	for r < len(s) {
		chara := s[r]
		index, hit := currMap[chara]
		if hit {
			// is the index past our current l 
			if index >= l {
				// move l up 
				l = index + 1
			}
		}

		currMap[chara] = r
		longest = max(longest, r - l + 1)
		r++
	}

	return longest
}
