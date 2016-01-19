package taxiFare

type calFare func(float64) float64

func DisCal(d float64) float64{
	var f float64
	if d <= flagDownDis {
		f = flagDownFare
	}
	if d > flagDownDis && d < fareIncreaseDis {
		f = (flagDownFare + (d - flagDownDis) * farePerKilometer)
	}
	if d >= fareIncreaseDis {
		f = (flagDownFare + (fareIncreaseDis - flagDownDis) * farePerKilometer + (d - fareIncreaseDis) * farePerKilometer * (1 + fareIncreasePercent))
	}
	return f
}

func TimeCal(t float64) float64{
	return (t * farePerMinute)
}

func TotalCal(d float64, t float64, DisFn calFare, TimeFn calFare) float64{
	return RoundPlus(DisFn(d) + TimeFn(t), 2)
}