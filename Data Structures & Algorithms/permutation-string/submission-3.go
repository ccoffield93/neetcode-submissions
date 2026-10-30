import "maps"
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[byte]int)
	for _, letter := range s1 {
		s1Map[byte(letter)]++
	}

	// starting window 
	testWindowMap := make(map[byte]int) 
	for i:=0; i < len(s1); i++ {
		testWindowMap[byte(s2[i])]++

	}
	l:= 0
	r:= len(s1)
	for {
		if maps.Equal(s1Map, testWindowMap) {
			return true
		} else {

			// If we've reached the end of s2, stop
			if r >= len(s2) {
				break
			}
			// add next, remove first 
			testWindowMap[s2[r]]++
			testWindowMap[s2[l]]--
			count, _ := testWindowMap[s2[l]]
			if count == 0 {
				delete(testWindowMap, byte(s2[l]))
			}
		}
		l++
		r++
	}
	return false
}
