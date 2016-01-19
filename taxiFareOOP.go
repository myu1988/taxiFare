package taxiFare

type TaxiFareOOP struct{
	Distance, Time float64
}

func (t TaxiFareOOP) DisCal() float64{
	var f float64
	if t.Distance <= flagDownDis {
		f = flagDownFare
	}
	if t.Distance > flagDownDis && t.Distance < fareIncreaseDis {
		f = (flagDownFare + (t.Distance - flagDownDis) * farePerKilometer)
	}
	if t.Distance >= fareIncreaseDis {
		f = (flagDownFare + (fareIncreaseDis - flagDownDis) * farePerKilometer + (t.Distance - fareIncreaseDis) * farePerKilometer * (1 + fareIncreasePercent))
	}
	return f
}

func (t TaxiFareOOP) TimeCal() float64{
	return (t.Time * farePerMinute)
}

func (t TaxiFareOOP) TotalCal() float64{
	return RoundPlus(t.DisCal() + t.TimeCal(), 2)
}