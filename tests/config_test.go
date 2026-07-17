package tests

import (
	"testing"

	"calculator-backend/internal/config"
	"github.com/stretchr/testify/require"
)

func TestConfigLoadUsesDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("GIN_MODE", "")
	t.Setenv("LOG_LEVEL", "")

	cfg := config.Load()

	require.Equal(t, "8080", cfg.Port)
	require.Equal(t, "debug", cfg.GinMode)
	require.Equal(t, "info", cfg.LogLevel)
}

func TestConfigLoadUsesEnvironmentValues(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("GIN_MODE", "release")
	t.Setenv("LOG_LEVEL", "warn")

	cfg := config.Load()

	require.Equal(t, "9090", cfg.Port)
	require.Equal(t, "release", cfg.GinMode)
	require.Equal(t, "warn", cfg.LogLevel)
}
