package taxiFare

import "testing"

func TestTaxiFareFun(t *testing.T) {
	fare := TotalCal(2,2,DisCal,TimeCal)
	if fare != 6.5 {
		t.Errorf("incorrect")
	}
	fare = TotalCal(4,4,DisCal,TimeCal)
	if fare != 10 {
		t.Errorf("incorrect",fare)
	}
	fare = TotalCal(10,4,DisCal,TimeCal)
	if fare != 20.5 {
		t.Errorf("incorrect",fare)
	}
}