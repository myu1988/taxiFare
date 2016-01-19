package taxiFare

import "testing"

func TestTaxiFareChan(t *testing.T) {
	c := make(chan float64)
	go TotalCalChan(2,2,c)
	fare := <- c
	if fare != 6.5 {
		t.Errorf("incorrect",fare)
	}
	go TotalCalChan(4,4,c)
	fare = <- c
	if fare != 10 {
		t.Errorf("incorrect",fare)
	}
	go TotalCalChan(10,4,c)
	fare = <- c
	if fare != 20.5 {
		t.Errorf("incorrect",fare)
	}
}