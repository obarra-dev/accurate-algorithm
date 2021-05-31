package alggo

import "testing"

var caseTests = []struct {
	word   string
	expect bool
}{
	{"(c(oder)) b(yte)", true},
	{"((c(oder)) ()()()(((omar))) b(yt(test)e))", true},
	{"", true},
	{"()", true},
	{"(coder)(byte))", false},
	{"(((((coder)(byte))", false},
	{"((((()", false},
}

func Test_bracketMatcher(t *testing.T) {
	for i, ct := range caseTests {
		got := bracketMatcher(ct.word)
		if got != ct.expect {
			t.Errorf("Test[%d]: bracketMatcher(%s) expect %v, got %v",
				i, ct.word, ct.expect, got)
		}
	}
}

func Test_bracketMatcherWithCounter(t *testing.T) {
	for i, ct := range caseTests {
		got := bracketMatcherWithCounter(ct.word)
		if got != ct.expect {
			t.Errorf("Test[%d]: bracketMatcher(%s) expect %v, got %v",
				i, ct.word, ct.expect, got)
		}
	}
}
