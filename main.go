package main

// Embedding allows types to be open for extension
// A is embedded into B, as both structs are in the same package,
// B has access to A's private fields
// We're extending A
import (
	"ei/ei"
	"fmt"
)

type B struct {
	ei.A
}

func (b B) Greet() { fmt.Println("Bienvenido Mundo ", b.Year) }

func main() {
	var a ei.A
	a.Year = 2016
	var b B
	b.Year = 2016
	a.Greet()
	b.Greet()
}
