package alggo

import (
	"fmt"
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

func Test_findFirstDuplicated(t *testing.T) {
	var caseTest = []struct {
		arg    []int
		expect int
	}{
		{[]int{1, 2, 4, 5, 5}, 5},
		{[]int{1, 2, 4, 2, 5, 5}, 2},
		{[]int{1, 2, 4, 5, 5, 2}, 5},
		{[]int{1, 2, 4, 3, 5, 6}, 0},
	}

	for i, ct := range caseTest {
		got := findFirstDuplicated(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: findFirstDuplicated(%v) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}

func Test_calculateSqrtSorted(t *testing.T) {
	var caseTest = []struct {
		arg    []float64
		expect []float64
	}{
		{[]float64{1, 2, 4, 5, 5}, []float64{2, 4, 16, 25, 25}},
		{[]float64{-2, -1, 4, 5, 5}, []float64{1, 4, 16, 25, 25}},
	}
	fmt.Println("omar rules")

	for i, ct := range caseTest {
		calculateSqrtSorted(ct.arg)

		fmt.Println("Input:", i)

	}
}

func Test_jumpingOnClouds(t *testing.T) {
	var caseTest = []struct {
		arg    []int
		expect int
	}{
		{[]int{0, 0, 1, 0, 0}, 3},
		{[]int{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int{0, 0, 0, 0, 1, 0}, 3},
		{[]int{0, 0, 0, 0, 0, 0}, 3},
	}

	for i, ct := range caseTest {
		got := jumpingOnClouds(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: jumpingOnClouds(%v) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}

func Test_salesByMatch(t *testing.T) {
	var caseTest = []struct {
		arg    []int
		expect int
	}{
		{[]int{0, 0, 1, 0, 0}, 2},
		{[]int{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	}

	for i, ct := range caseTest {
		got := salesByMatch(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: salesByMatch(%v) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}

func Test_detectNetwork(t *testing.T) {
	var caseTest = []struct {
		arg    string
		expect string
	}{
		{"341234567890123", "American Express"},
		{"371234567890123", "American Express"},
		{"38123456789012", "Diners Club"},
		{"5112345678901245", "MasterCard"},
		{"5212345678901245", "MasterCard"},
		{"5312345678901245", "MasterCard"},
		{"5412345678901245", "MasterCard"},
		{"5512345678901245", "MasterCard"},
		{"4512345678901", "Visa"},
		{"4512345678901345", "Visa"},
		{"4512345678901345345", "Visa"},
	}

	for i, ct := range caseTest {
		got := detectNetwork(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: detectNetwork(%v) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}
