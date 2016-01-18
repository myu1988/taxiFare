package taxiFare

type calFare func(float64) float64

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

func TotalCal(d float64, t float64, DisFn calFare, TimeFn calFare) float64{
	return RoundPlus(DisFn(d) + TimeFn(t), 2)
}