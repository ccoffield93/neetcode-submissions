func checkInclusion(s1 string, s2 string) bool {
	contents := make(map[rune]int)

	for _, char := range s1 {
		val, found := contents[char]
		{
			if found {
				contents[char] = val + 1
			} else {
				contents[char] = 1
			}
		}
	}

	left := 0
	right := len(s1) - 1
	for right < len(s2) {
		// optimization-- if s2[right] not in s1, move on
		/*
		_, skipIfTrue := contents[rune(s2[right])]
		if skipIfTrue{
			left++
			right++
			continue
		}*/

		checkMap := make(map[rune]int)
		for i:= left; i <= right; i++ {
			letter := rune(s2[i])
			_, found := contents[letter]
			if !found {
				break; // one of the letters isn't matching
			} else {
				curr, found := checkMap[letter]
				if found {
					checkMap[letter] = curr + 1
				} else {
					checkMap[letter] = 1
				}
			}
		}

		mismatch := false 
		for key, value := range contents {
			matchFreq, found := checkMap[key]
			if found {
				if matchFreq != value {
					mismatch = true
				}
			} else {
				mismatch = true
			}
		}

		if !mismatch {
			return true
		}


		left++
		right++
	}

	return false
}
