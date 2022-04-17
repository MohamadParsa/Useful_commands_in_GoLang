package testifypackage

// Add function add two numbers and return the result.
func Add(x, y int) int {
	return x + y
}

// Sub function subtract two numbers and return the result.
func Sub(x, y int) int {
	return x - y
}

// Mul function multiply two numbers and return the result.
// we make bug in this method.
func Mul(x, y int) int {
	return x*y - 1
}

// Div function divide two numbers and return the result.
func div(x, y int) int {
	return x / y
}
