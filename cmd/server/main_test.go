package main

import (
	"errors"
	"net/http"
	"testing"

	"calculator-backend/internal/config"
	"github.com/stretchr/testify/require"
)

func TestBuildLoggerAndRouter(t *testing.T) {
	cfg := config.Config{Port: "8080", GinMode: "release", LogLevel: "debug"}
	logger := buildLogger(cfg)
	require.NotNil(t, logger)

	router := buildRouter()
	require.NotNil(t, router)

	var handler http.Handler = router
	require.NotNil(t, handler)
}

func TestRunReturnsZeroOnSuccessfulStartup(t *testing.T) {
	t.Setenv("PORT", "8080")
	originalListenAndServe := listenAndServe
	defer func() { listenAndServe = originalListenAndServe }()

	listenAndServe = func(address string, handler http.Handler) error {
		require.Equal(t, ":8080", address)
		require.NotNil(t, handler)
		return nil
	}

	code := run()
	require.Equal(t, 0, code)
}

func TestRunReturnsOneWhenStartupFails(t *testing.T) {
	t.Setenv("PORT", "9090")
	originalListenAndServe := listenAndServe
	defer func() { listenAndServe = originalListenAndServe }()

	listenAndServe = func(address string, handler http.Handler) error {
		require.Equal(t, ":9090", address)
		require.NotNil(t, handler)
		return errors.New("boom")
	}

	code := run()
	require.Equal(t, 1, code)
}

func TestMainCallsExitWithRunCode(t *testing.T) {
	t.Setenv("PORT", "8080")
	originalListenAndServe := listenAndServe
	originalOsExit := osExit
	defer func() {
		listenAndServe = originalListenAndServe
		osExit = originalOsExit
	}()

	listenAndServe = func(address string, handler http.Handler) error {
		require.Equal(t, ":8080", address)
		require.NotNil(t, handler)
		return nil
	}

	var code int
	osExit = func(exitCode int) {
		code = exitCode
		panic(exitCode)
	}

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("expected main to exit")
		}
	}()

	main()
	require.Equal(t, 0, code)
}
