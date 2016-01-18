package taxiFare

import "testing"

func TestTaxiFareFun(t *testing.T) {
	fare := TotalCal(2,2,DisCal,TimeCal)
	if fare != 6.5 {
		t.Errorf("incorrect")
	}
}