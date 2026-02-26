package calculator

import (
	"errors"
)

// Add soma dois números
func Add(a, b float64) float64 {
	return a + b
}

// Subtract subtrai o segundo número do primeiro
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplica dois números
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divide o primeiro número pelo segundo
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil
}
