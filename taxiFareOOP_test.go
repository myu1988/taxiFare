package taxiFare

import "testing"

func TestTaxiFareOOP(t *testing.T) {
	fare := TaxiFareOOP{2,2}.TotalCal()
	if fare != 6.5 {
		t.Errorf("incorrect",fare)
	}
	fare = TaxiFareOOP{4,4}.TotalCal()
	if fare != 10 {
		t.Errorf("incorrect",fare)
	}
	fare = TaxiFareOOP{10,4}.TotalCal()
	if fare != 20.5 {
		t.Errorf("incorrect",fare)
	}
}