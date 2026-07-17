package tests

import (
	"os"
	"testing"

	"calculator-backend/internal/config"
	"calculator-backend/internal/evaluator"

	"github.com/stretchr/testify/require"
)

func TestEvaluatorAndConfig(t *testing.T) {
	t.Run("evaluate arithmetic and functions", func(t *testing.T) {
		eval := evaluator.NewEvaluator()
		result, err := eval.Evaluate("sqrt(9)+2")
		require.NoError(t, err)
		require.Equal(t, 5.0, result)
	})

	t.Run("evaluate division by zero", func(t *testing.T) {
		eval := evaluator.NewEvaluator()
		_, err := eval.Evaluate("1/0")
		require.Error(t, err)
	})

	t.Run("load config from environment", func(t *testing.T) {
		require.NoError(t, os.Setenv("PORT", "9090"))
		require.NoError(t, os.Setenv("GIN_MODE", "release"))
		require.NoError(t, os.Setenv("LOG_LEVEL", "warn"))
		defer os.Unsetenv("PORT")
		defer os.Unsetenv("GIN_MODE")
		defer os.Unsetenv("LOG_LEVEL")

		loaded := config.Load()
		require.Equal(t, "9090", loaded.Port)
		require.Equal(t, "release", loaded.GinMode)
		require.Equal(t, "warn", loaded.LogLevel)
	})
}
