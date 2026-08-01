func characterReplacement(s string, k int) int {
	left := 0
	right := 0

	longestLen := 0

	mapOfContents := make(map[byte]int)
	for (right < len(s)) {
		val := s[right]
		current, found := mapOfContents[val]
		if found {
			mapOfContents[val] = current + 1
		} else {
			mapOfContents[val] = 1
		}

		// the number of characters NOT the most frequent 
		// must
		_, freq := getMostFrequentKey(mapOfContents) 
		for ((right - left - freq) >= k) {
			// we have to move left pointer 
			losing := s[left]
			mapOfContents[losing]--
			left++
			_, freq = getMostFrequentKey(mapOfContents)
		}

		longestLen = max(longestLen, right - left + 1)

		right++
	}

	return longestLen
}

func getMostFrequentKey(in map[byte]int) (byte, int) {
	highestFreq := 0
	var mostFreqKey byte

	for key, value := range(in) {
		if value > highestFreq {
			highestFreq = value
			mostFreqKey = key
		}
	}

	return mostFreqKey, highestFreq
}
