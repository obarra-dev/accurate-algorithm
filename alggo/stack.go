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

// https://leetcode.com/problems/remove-k-digits/
func removeKDigits(digits string, k int) string {
	if k >= len(digits) {
		return "0"
	}

	var stack []rune
	for _, v := range digits {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > v {
			//pop
			stack = stack[:len(stack)-1]
			k--
		}

		if len(stack) > 0 || v != '0' {
			stack = append(stack, v)
		}
	}

	for len(stack) > 0 && k > 0 {
		//pop
		stack = stack[:len(stack)-1]
		k--
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}

func countStudentsWithoutLunch(lunches []uint, preferences []uint) uint {
	//lunches is a inverted stack
	//preferences is a queue
	var elements = uint(len(preferences))
	for elements > 0 {
		top := lunches[0]
		tail := preferences[0]
		//dequeue
		preferences = preferences[1:]
		if tail == top {
			//pop
			lunches = lunches[1:]
			elements = uint(len(preferences))
		} else {
			//enqueue
			preferences = append(preferences, tail)
			elements--
		}
	}

	return uint(len(preferences))
}
