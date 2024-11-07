package base

import "testing"

func Test_primeCount(t *testing.T) {
	c := primeCount1(100)
	if c != 25 {
		t.Error("primeCount 100 error")
	}
	c2 := primeCount2(100)
	if c2 != 25 {
		t.Error("primeCount 100 error")
	}
}
