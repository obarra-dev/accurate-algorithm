package alggo

import "testing"

func Test_bracketMatcher(t *testing.T) {
	caseTests := []struct {
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
	for i, ct := range caseTests {
		got := bracketMatcher(ct.word)
		if got != ct.expect {
			t.Errorf("Test[%d]: bracketMatcher(%s) expect %v, got %v",
				i, ct.word, ct.expect, got)
		}
	}
}
