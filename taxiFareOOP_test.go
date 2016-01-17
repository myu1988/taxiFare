package taxiFare

import "testing"

func TestTaxiFareOOP(t *testing.T) {
	fare := TaxiFareOOP{2,2}
	if fare.TotalCal() != 6.5 {
		t.Errorf("incorrect")
	}
}