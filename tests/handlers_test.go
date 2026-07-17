package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-backend/internal/api"
	"calculator-backend/internal/evaluator"
	"calculator-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := service.NewCalculatorService(evaluator.NewEvaluator())
	router := api.SetupRoutes(svc)

	t.Run("calculate success", func(t *testing.T) {
		body := bytes.NewBufferString(`{"expression":"2+2*2"}`)
		req := httptest.NewRequest(http.MethodPost, "/calculate", body)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)
		require.JSONEq(t, `{"result":6}`, recorder.Body.String())
	})

	t.Run("calculate invalid expression", func(t *testing.T) {
		body := bytes.NewBufferString(`{"expression":"2+"}`)
		req := httptest.NewRequest(http.MethodPost, "/calculate", body)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("calculate zero result", func(t *testing.T) {
		body := bytes.NewBufferString(`{"expression":"2*(2-2)"}`)
		req := httptest.NewRequest(http.MethodPost, "/calculate", body)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)
		require.JSONEq(t, `{"result":0}`, recorder.Body.String())
	})

	t.Run("health", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)
		require.JSONEq(t, `{"status":"ok"}`, recorder.Body.String())
	})
}
