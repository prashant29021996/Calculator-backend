package tests

import (
	"testing"

	"calculator-backend/internal/validator"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantErr    bool
	}{
		{name: "valid arithmetic", expression: "2+2*2", wantErr: false},
		{name: "valid parentheses", expression: "(2+5)*8", wantErr: false},
		{name: "valid function", expression: "sqrt(25)", wantErr: false},
		{name: "valid uppercase function", expression: "SQRT(25)", wantErr: false},
		{name: "empty expression", expression: "", wantErr: true},
		{name: "invalid character", expression: "2 & 3", wantErr: true},
		{name: "malformed expression", expression: "2+", wantErr: true},
		{name: "unbalanced parentheses", expression: "(2+3", wantErr: true},
		{name: "unsupported function", expression: "foo(2)", wantErr: true},
		{name: "division by zero", expression: "10/0", wantErr: true},
		{name: "expression too long", expression: string(make([]byte, 257)), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.expression)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
