package Good_example

type SubstractCalculator struct {
	Calculator
}

func (calc *Calculator) Substract(a float64, b float64) {
	calc.result = a - b
}
