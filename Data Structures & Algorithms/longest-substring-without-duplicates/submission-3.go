func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s) // 0 or 1
	}
	indexMap := make(map[byte]int)
	left := 0
	right := 0
	longest := 1

	n := len(s) 
	for (right < n) {
		current := s[right]
	
        // If current character was seen and is within the active window
        if idx, hit := indexMap[current]; hit {
            left = max(left, idx+1)
        }
    
		// add current value, check longest, move on 
		indexMap[s[right]] = right
		longest = max(longest, right - left + 1)
		right++
	}

	return longest
}
