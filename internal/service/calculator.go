package service

import (
	"calculator-backend/internal/validator"
)

type Evaluator interface {
	Evaluate(expression string) (float64, error)
}

type CalculatorService interface {
	Calculate(expression string) (float64, error)
}

type calculatorService struct {
	evaluator Evaluator
}

func NewCalculatorService(evaluator Evaluator) CalculatorService {
	return &calculatorService{evaluator: evaluator}
}

func (s *calculatorService) Calculate(expression string) (float64, error) {
	if err := validator.Validate(expression); err != nil {
		return 0, err
	}

	result, err := s.evaluator.Evaluate(expression)
	if err != nil {
		return 0, err
	}
	return result, nil
}
