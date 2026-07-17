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

func TestAPIContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := service.NewCalculatorService(evaluator.NewEvaluator())
	router := api.SetupRoutes(svc)

	t.Run("rejects malformed json body", func(t *testing.T) {
		body := bytes.NewBufferString(`{"expression":`)
		req := httptest.NewRequest(http.MethodPost, "/calculate", body)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.JSONEq(t, `{"error":"invalid request body","result":0}`, recorder.Body.String())
	})

	t.Run("returns validation error for unsupported function", func(t *testing.T) {
		body := bytes.NewBufferString(`{"expression":"foo(2)"}`)
		req := httptest.NewRequest(http.MethodPost, "/calculate", body)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.Contains(t, recorder.Body.String(), "invalid function")
	})
}
