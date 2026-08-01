func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s) // 0 or 1
	}
	indexMap := make(map[byte]int)
	left := 0
	right := 0
	longest := 1

	for right < len(s) {
		current := s[right]
		lastIndex, found := indexMap[current]
		if found && lastIndex >= left {
			// we now have a duplicate 
			for left < right {
				toRemove := s[left]
				delete(indexMap, toRemove)
				left++ 
				if toRemove == current { 
					// we have removed the item that 'right' is pointing to
					indexMap[toRemove] = right
					break
				}
			}

		} else {
			// no duplicate, check if max and keep going
			longest = max(longest, right - left + 1)
			indexMap[current] = right
		}
		right++
	}

	return longest
}
