package alggo

import (
	"testing"
)

var caseTestFlatten = []struct {
	arg   []interface{}
	expect string
}{
	{[]interface{}{1, []interface{} {1, 2},[]interface{} {[]interface{} {3, 4}}, []interface{} {1, []interface{}{}}},
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
