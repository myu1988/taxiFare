package taxiFare

type TaxiFareOOP struct{
	Distance, Time float64
}

func (t TaxiFareOOP) DisCal() float64{
	return DisCal(t.Distance)
}

func (t TaxiFareOOP) TimeCal() float64{
	return TimeCal(t.Time)
}

func (t TaxiFareOOP) TotalCal() float64{
	return RoundPlus(t.DisCal() + t.TimeCal(), 2)
}