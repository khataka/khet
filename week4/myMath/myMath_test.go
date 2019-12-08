package myMath

import "testing"

func TestAverage(t *testing.T) {
	v := Average(3, 5)
	if v != 4 {
		t.Error("Avarage of 3,5 is 4, not", v)
	}

}
