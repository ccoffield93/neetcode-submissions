

func isValid(s string) bool {
    openRound := 0
	openCurly := 0
	openSquare := 0
	stack := []byte{}

	for _, char := range s {
		switch(char) {
			// adding to the stack 
			case '(': 
				stack = append(stack, '(')
				openRound++
			case '[':
				stack = append(stack, '[')
				openSquare++
			case '{':
				stack = append(stack, '{')
				openCurly++
			// closing out a bracket type
			case ')':
				// need to check that the last item pushed WAS the match
				if len(stack) > 0 && stack[len(stack) - 1] == '(' {
					openRound--
					stack = stack[:len(stack)-1]
				} else {
					return false 
				}
			case ']':
				// need to check that the last item pushed WAS the match
				if len(stack) > 0 && stack[len(stack) - 1] == '[' {
					openSquare--
					stack = stack[:len(stack)-1]
				} else {
					return false 
				}
			case '}':
				// need to check that the last item pushed WAS the match
				if len(stack) > 0 && stack[len(stack) - 1] == '{' {
					openCurly--
					stack = stack[:len(stack)-1]
				} else {
					return false 
				}
			default: 
				return false
		}
	}

	if openRound == 0 && openCurly == 0 && openSquare == 0 {
		return true
	} 

	return false
}
