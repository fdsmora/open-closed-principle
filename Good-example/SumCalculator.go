package Good_example

type SumCalculator struct {
	Calculator
}

func (calc *Calculator) Sum(a float64, b float64) {
	calc.result = a + b
}
