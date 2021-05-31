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

func bracketMatcherWithCounter(word string) bool {
	var counter int
	for _, s := range word {
		switch {
		case s == '(':
			counter++
		case s == ')':
			counter--
		case counter < 0:
			return false
		}
	}

	if counter == 0 {
		return true
	}

	return false
}
