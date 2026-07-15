func isPalindrome(s string) bool {
  // remove everything that is not alphanumeric and make it case-insensitive
  s = removeNonAlphanumericFast(s)
  s = strings.ToLower(s)
  i := 0
  j := len(s) - 1

  // go until they touch or cross, touching in the middle doesn't need a check
  for i < j {
	a := s[i]
	b := s[j]
	if a != b {
		return false
	}
	i++
	j--
  }
	return true
}

func removeNonAlphanumericFast(str string) string {
	var sb strings.Builder
	sb.Grow(len(str)) // Optimize memory allocation
	
	for i := 0; i < len(str); i++ {
		b := str[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			sb.WriteByte(b)
		}
	}
	return sb.String()
}