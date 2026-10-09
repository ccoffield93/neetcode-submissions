func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	sStripped := ""
	for i := 0; i < len(s); i++ {
		chara := s[i]
		if ('a' <= chara && chara <= 'z') || ('0' <= chara && chara <= '9') {
			sStripped += string(chara)
		}
	}

	l := 0
	r := len(sStripped) - 1
	for l <= r {
		a := sStripped[l]
		b := sStripped[r]
		if a != b {
			return false
		}
		l++
		r--
	}

	return true
}
