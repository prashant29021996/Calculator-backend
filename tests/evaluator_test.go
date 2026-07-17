package tests

import (
	"testing"

	"calculator-backend/internal/evaluator"
	"github.com/stretchr/testify/require"
)

func TestEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{name: "arithmetic", expression: "2+2*2", want: 6},
		{name: "parentheses", expression: "(2+5)*8", want: 56},
		{name: "function", expression: "sqrt(25)", want: 5},
		{name: "division by zero", expression: "10/0", wantErr: true},
		{name: "malformed expression", expression: "2+", wantErr: true},
		{name: "unsupported function", expression: "foo(2)", wantErr: true},
		{name: "sqrt of negative", expression: "sqrt(-1)", wantErr: true},
		{name: "sqrt missing argument", expression: "sqrt()", wantErr: true},
	}

	engine := evaluator.NewEvaluator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := engine.Evaluate(tt.expression)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.InDelta(t, tt.want, result, 0.0001)
		})
	}
}
