package alggo

func bracketMatcher(word string) bool {
	var stack []int32
	for _, s := range word {
		if s == '(' {
			stack = append(stack, '*')
		} else if s == ')' {
			// pop
			n := len(stack)
			if n == 0 {
				return false
			}
			n--
			stack = stack[:n]
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
