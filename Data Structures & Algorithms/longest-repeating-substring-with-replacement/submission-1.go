func characterReplacement(s string, k int) int {

	mapOfFreq := make(map[byte]int)

	longest := 0 
	if len(s) <= k {
		return len(s)
	}

	l := 0
	r := 0 
	for r < len(s) {
		// we can replace up to K characters, so the longest substring 
		// must only have that many that AREN'T the most frequent 
		mapOfFreq[s[r]]++

		mostFreq := highestValInMap(mapOfFreq)
		for (r - l - mostFreq + 1 > k) {
			mapOfFreq[s[l]]--
			l++
			mostFreq = highestValInMap(mapOfFreq)
		}
		longest = max(longest, r - l + 1)
	    r++
	}

	return longest
}

func highestValInMap(check map[byte]int) int {
	if len(check) == 0 {
		return 0
	}
	highest := -1
	for _, val := range(check) {
		highest = max(highest, val)
	}
	return highest
}