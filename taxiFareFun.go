package taxiFare

func DisCal(d float64) float64{
	var f float64
	if d <= 2 {
		f = 6
	}
	if d > 2 && d < 8 {
		f = (6 + (d - 2) * 1.5)
	}
	if d > 8 {
		f = (6 + 6 * 1.5 + (d - 8) * 1.5 * 1.5)
	}
	return f
}

func TimeCal(t float64) float64{
	return (t * 0.25)
}

func TotalCal(d, t float64) float64{
	return RoundPlus(DisCal(d) + TimeCal(t), 2)
}