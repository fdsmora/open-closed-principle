package Calculator

type Calculator struct {
	result float64
}

func (calc *Calculator) Result() float64 {
	return calc.result
}
