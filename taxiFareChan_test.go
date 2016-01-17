package taxiFare
import "fmt"
import "testing"

func TestTaxiFareChan(t *testing.T) {
	c := make(chan float64)
	go TotalCalChan(2,2,c)
	fare := <- c
	fmt.Println(fare)
	if fare != 6.5 {
		t.Errorf("incorrect")
	}
}