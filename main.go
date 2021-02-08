package main

import (
	Good_example "Calculator/Good-example"
	"fmt"
)

func main() {
	good_example()
}

func good_example() {
	var sumCalc Good_example.SumCalculator
	var substractCalc Good_example.SubstractCalculator

	sumCalc.Sum(2, 3)
	fmt.Println("Sum is ", sumCalc.Result())
	substractCalc.Substract(6, 2)
	fmt.Println("Substraction is ", substractCalc.Result())
}
