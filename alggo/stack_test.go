package alggo

import "testing"

var caseTestsBracketMatcher = []struct {
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
	for i, ct := range caseTestsBracketMatcher {
		got := bracketMatcher(ct.word)
		if got != ct.expect {
			t.Errorf("Test[%d]: bracketMatcher(%s) expect %v, got %v",
				i, ct.word, ct.expect, got)
		}
	}
}

func Test_bracketMatcherWithCounter(t *testing.T) {
	for i, ct := range caseTestsBracketMatcher {
		got := bracketMatcherWithCounter(ct.word)
		if got != ct.expect {
			t.Errorf("Test[%d]: bracketMatcher(%s) expect %v, got %v",
				i, ct.word, ct.expect, got)
		}
	}
}

func Test_removeKDigits(t *testing.T) {
	var caseTest = []struct {
		digits string
		k      int
		expect string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"1230987", 2, "10987"},
		{"1230987", 3, "987"},
		{"100111", 2, "11"},
		{"1001", 2, "0"},
		{"1001111222", 4, "1111"},
	}

	for i, ct := range caseTest {
		got := removeKDigits(ct.digits, ct.k)
		if got != ct.expect {
			t.Errorf("Test[%d]: removeKDigits(%v, %v) expect %v, got %v",
				i, ct.digits, ct.k, ct.expect, got)
		}
	}
}

func Test_countStudentsWithoutLunch(t *testing.T) {
	var caseTest = []struct {
		typeLunches []uint
		preferences []uint
		expect      uint
	}{
		{[]uint{0, 1, 0, 0}, []uint{0, 0, 0, 0}, 3},
		{[]uint{0, 0, 1, 0}, []uint{1, 0, 0, 0}, 0},
		{[]uint{0, 1, 1, 0, 1, 0}, []uint{1, 1, 1, 0, 1, 0}, 1},
	}

	for i, ct := range caseTest {
		got := countStudentsWithoutLunch(ct.typeLunches, ct.preferences)
		if got != ct.expect {
			t.Errorf("Test[%d]: countStudentsWithoutLunch(%v, %v) expect %v, got %v",
				i, ct.typeLunches, ct.preferences, ct.expect, got)
		}
	}
}
