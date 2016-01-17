package taxiFare

type TaxiFareOOP struct{
	Distance, Time float64
}
func (t TaxiFareOOP) DisCal() float64{
	var f float64
	if t.Distance<=2 {
		f = 6
	}
	if t.Distance>2 && t.Distance<8 {
		f = (6 + (t.Distance - 2) * 1.5)
	}
	if t.Distance > 8 {
		f = (6 + 6 * 1.5 + (t.Distance - 8) * 1.5 * 1.5)
	}
	return f
}

func (t TaxiFareOOP) TimeCal() float64{
	return (t.Time * 0.25)
}

func (t TaxiFareOOP) TotalCal() float64{
	return RoundPlus(t.DisCal() + t.TimeCal(), 2)
}