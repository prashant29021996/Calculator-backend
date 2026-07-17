package tests

import (
	"errors"
	"testing"

	"calculator-backend/internal/service"
	"github.com/stretchr/testify/require"
)

type stubEvaluator struct {
	result float64
	err    error
	expr   string
}

func (s *stubEvaluator) Evaluate(expression string) (float64, error) {
	s.expr = expression
	return s.result, s.err
}

func TestCalculatorService_Calculate(t *testing.T) {
	evaluator := &stubEvaluator{result: 6}
	svc := service.NewCalculatorService(evaluator)

	result, err := svc.Calculate("2+2*2")
	require.NoError(t, err)
	require.Equal(t, 6.0, result)
	require.Equal(t, "2+2*2", evaluator.expr)
}

func TestCalculatorService_CalculateError(t *testing.T) {
	evaluator := &stubEvaluator{err: errors.New("division by zero")}
	svc := service.NewCalculatorService(evaluator)

	_, err := svc.Calculate("10/0")
	require.Error(t, err)
	require.Equal(t, "division by zero", err.Error())
}
