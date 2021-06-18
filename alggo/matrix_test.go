package alggo

import "testing"

func Test_highestHourglass(t *testing.T) {
	var caseTest = []struct {
		arg    [6][6]int
		expect int
	}{
		{[6][6]int{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}}, 7},
		{[6][6]int{{-9, -9, -9, 1, 1, 1}, {0, -9, 0, 4, 3, 2}, {-9, -9, -9, 1, 2, 3}, {0, 0, 8, 6, 6, 0}, {0, 0, 0, -2, 0, 0}, {0, 0, 1, 2, 4, 0}}, 28},
		{[6][6]int{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 0, 2, 4, 4, 0}, {0, 0, 0, 2, 0, 0}, {0, 0, 1, 2, 4, 0}}, 19},
	}

	for i, ct := range caseTest {
		got := highestHourglass(ct.arg)
		if got != ct.expect {
			t.Errorf("Test[%d]: highestHourglass(%v) expect %v, got %v",
				i, ct.arg, ct.expect, got)
		}
	}
}
