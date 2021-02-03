package main

import (
	"Calculator/Calculator"
	"fmt"
)

func main() {

	var sumCalc Calculator.SumCalculator
	var substractCalc Calculator.SubstractCalculator

	sumCalc.Sum(2, 3)
	fmt.Println("Sum is ", sumCalc.Result())
	substractCalc.Substract(6, 2)
	fmt.Println("Substraction is ", sumCalc.Result())
}
