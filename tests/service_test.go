package tests

import (
	"testing"

	"calculator-backend/internal/evaluator"
	"calculator-backend/internal/service"

	"github.com/stretchr/testify/require"
)

func TestCalculatorService(t *testing.T) {
	calculator := service.NewCalculatorService(evaluator.NewEvaluator())

	t.Run("calculate valid expression", func(t *testing.T) {
		result, err := calculator.Calculate("2+2*2")
		require.NoError(t, err)
		require.Equal(t, 6.0, result)
	})

	t.Run("reject invalid expression", func(t *testing.T) {
		_, err := calculator.Calculate("2+")
		require.Error(t, err)
	})
}
