// uc_test.go
package uc

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		u := UpperCase(ut.in)
		if u != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s.", ut.in, u, ut.out)
		}
	}
}
