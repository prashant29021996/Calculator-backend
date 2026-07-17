package evaluator

import (
	"fmt"
	"math"
	"strings"

	"github.com/Knetic/govaluate"
)

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Evaluate(expression string) (float64, error) {
	expression = strings.TrimSpace(expression)
	if expression == "" {
		return 0, fmt.Errorf("expression is required")
	}

	expression = strings.ReplaceAll(expression, "sqrt", "SQRT")
	parameters := make(map[string]interface{})
	expression = strings.ToUpper(expression)

	expr, err := govaluate.NewEvaluableExpressionWithFunctions(expression, map[string]govaluate.ExpressionFunction{
		"SQRT": func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("sqrt expects one argument")
			}
			value, ok := args[0].(float64)
			if !ok {
				return nil, fmt.Errorf("sqrt expects a numeric argument")
			}
			if value < 0 {
				return nil, fmt.Errorf("sqrt of negative number")
			}
			return math.Sqrt(value), nil
		},
	})
	if err != nil {
		return 0, err
	}

	result, err := expr.Evaluate(parameters)
	if err != nil {
		message := strings.ToLower(err.Error())
		if strings.Contains(message, "divide") || strings.Contains(message, "zero") || strings.Contains(message, "division") {
			return 0, fmt.Errorf("division by zero")
		}
		return 0, err
	}

	value, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected result type")
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("division by zero")
	}
	return value, nil
}
