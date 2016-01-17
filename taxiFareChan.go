package taxiFare

func TotalCalChan(d float64, t float64, c chan float64){
	c <- RoundPlus(DisCal(d) + TimeCal(t), 2)
}