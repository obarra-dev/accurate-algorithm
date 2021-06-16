package alggo

import (
	"testing"
)

var caseTestFlatten = []struct {
	arg    []interface{}
	expect string
}{
	{[]interface{}{1, []interface{}{1, 2}, []interface{}{[]interface{}{3, 4}}, []interface{}{1, []interface{}{}}},
		"[1,1,2,[3,4,]1,[]]"},
}

func Test_flattenLevelOne(t *testing.T) {
	for i, ct := range caseTestFlatten {
		got := flattenLevelOne(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: flattenLevelOne(%s) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}

var caseTestMostRepeated = []struct {
	arg    string
	expect string
}{
	{"omar alberto barra omar", "OMAR"},
	{"this a repeated word test this is a a", "A"},
}

func Test_mostRepeated(t *testing.T) {
	for i, ct := range caseTestMostRepeated {
		got := mostRepeated(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: mostRepeated(%s) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}


var caseTestIsPalindrome = []struct {
	arg    string
	expect bool
}{
	{"omar alberto barra omar", false},
	{"aaaaaaaaaa", true},
}

func Test_isPalindrome(t *testing.T) {
	for i, ct := range caseTestIsPalindrome {
		got := isPalindrome(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: isPalindrome(%s) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}
