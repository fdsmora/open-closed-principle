package Bad_example

type Calculator struct {
	result float64
}

func (calc *Calculator) Sum(a float64, b float64) {
	calc.result = a + b
}

func (calc *Calculator) Substract(a float64, b float64) {
	calc.result = a - b
}

func (calc *Calculator) Result() float64 {
	return calc.result
}
