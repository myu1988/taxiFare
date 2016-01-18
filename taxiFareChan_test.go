package taxiFare

import "testing"

func TestTaxiFareChan(t *testing.T) {
	c := make(chan float64)
	go TotalCalChan(2,2,c)
	fare := <- c
	if fare != 6.5 {
		t.Errorf("incorrect",fare)
	}
}