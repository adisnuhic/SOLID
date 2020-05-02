/*
	Definition:
	Software entities (classes, modules, functions, etc.) should be open for extension, but
	closed for modification. Essentially meaning that classes should be extended to change
	functionality, rather than being altered.
*/

package main

import "fmt"

/*--------------------------------------------------------*/
/*	     BREAKING OPEN/CLOSED PRINCIPLE	          */
/*--------------------------------------------------------*/
type CalculatorBad struct{}

// This Calculate function can ony do "+" operation against 2 numbers,
// but what if a want to extend Calculate to do "-" operation?
// Of course i will need to modify existing Calcuate function which will break
// this principle which says entities should be open fo extension and closed for modification
func (c CalculatorBad) Calculate(a int, b int) int {
	return a + b
}

// Commented Example below shows how we break this principle by modifying existing Calculate function
/*
func (c CalculatorBad) Calculate(a int, b int, operation string) int {
	if operation == "add" {
		return a + b
	}

	if operation == "minus" {
		return a - b
	}

	panic("Operation does not exists!")
}
*/

/*--------------------------------------------------------*/
/*	     IMPLEMENTING OPEN/CLOSED PRINCIPLE	          */
/*--------------------------------------------------------*/
// Instead i will create ICalculator interface which has Execute signature method
type ICalculator interface {
	Execute(int, int) int
}

// Add type implements ICalculator interface by implementing Execute method
type Add struct{}

func (Add) Execute(a int, b int) int {
	return a + b
}

// Minus type implements ICalculator interface by implementing Execute method
type Minus struct{}

func (Minus) Execute(a int, b int) int {
	return a - b
}

// CalculatorGood embeds ICalculator interface
type CalculatorGood struct {
	c ICalculator
}

func (c CalculatorGood) Execute(a int, b int) int {
	// c.c.Execute(a,b) will call Add->Execute() or Minus->Execute() depeding on what
	// type we pass. The reason for that is we can pass Add or Minus types through ICalculator
	// interface because both types implements this interface
	return c.c.Execute(a, b)
}

func main() {
	// Running bad example
	c1 := (&CalculatorBad{}).Calculate(2, 1)
	fmt.Printf("%v", c1)
	fmt.Println()

	// Running good example
	a := (&CalculatorGood{&Add{}}).Execute(5, 3)   // Passing Add type with data a=5, b=3
	b := (&CalculatorGood{&Minus{}}).Execute(5, 3) // Passing Minus type with data a=5, b=3
	fmt.Printf("%v", a)
	fmt.Println()
	fmt.Printf("%v", b)
	fmt.Println()
}
