package ei

import "fmt"

type A struct {
	Year int
}

func (a A) Greet() { fmt.Println("Hola mundo", a.Year) }
