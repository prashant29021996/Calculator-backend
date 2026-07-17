package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

const maxExpressionLength = 256

var validCharacters = regexp.MustCompile(`^[0-9+\-*/%^().\sA-Za-z]+$`)

func Validate(expression string) error {
	expr := strings.TrimSpace(expression)
	if expr == "" {
		return fmt.Errorf("expression is required")
	}
	if len(expr) > maxExpressionLength {
		return fmt.Errorf("expression is too long")
	}
	if !validCharacters.MatchString(expr) {
		return fmt.Errorf("invalid characters in expression")
	}

	if strings.Contains(expr, "/0") || strings.Contains(expr, "%0") {
		return fmt.Errorf("division by zero")
	}

	normalized := strings.ToUpper(expr)
	_, err := govaluate.NewEvaluableExpressionWithFunctions(normalized, map[string]govaluate.ExpressionFunction{
		"SQRT": func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("sqrt expects one argument")
			}
			return 0.0, nil
		},
	})
	if err != nil {
		message := strings.ToLower(err.Error())
		if strings.Contains(message, "function") {
			return fmt.Errorf("invalid function")
		}
		return fmt.Errorf("malformed expression")
	}
	return nil
}
