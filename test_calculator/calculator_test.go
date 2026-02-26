package calculator

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	// Teste de Soma
	t.Run("Soma", func(t *testing.T) {
		got := Add(10, 5)
		want := 15.0
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// Teste de Subtração
	t.Run("Subtração", func(t *testing.T) {
		got := Subtract(10, 5)
		want := 5.0
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// Teste de Multiplicação
	t.Run("Multiplicação", func(t *testing.T) {
		got := Multiply(10, 5)
		want := 50.0
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// Teste de Divisão
	t.Run("Divisão", func(t *testing.T) {
		t.Run("Divisão normal", func(t *testing.T) {
			got, err := Divide(10, 2)
			want := 5.0
			if err != nil {
				t.Errorf("erro inesperado: %v", err)
			}
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("Erro de divisão por zero", func(t *testing.T) {
			_, err := Divide(10, 0)
			if err == nil {
				t.Error("esperava um erro ao dividir por zero, mas não obteve")
			}
		})
	})
}
