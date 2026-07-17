package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-backend/internal/evaluator"
	"calculator-backend/internal/middleware"
	"calculator-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMiddlewareAndService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("request logger records status", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.RequestLogger())
		router.GET("/ping", func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})

		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("recovery converts panic to 500", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.Recovery())
		router.GET("/panic", func(c *gin.Context) {
			panic("boom")
		})

		req := httptest.NewRequest(http.MethodGet, "/panic", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		require.Equal(t, http.StatusInternalServerError, rec.Code)
		require.Contains(t, rec.Body.String(), "internal server error")
	})

	t.Run("calculator service returns evaluator result", func(t *testing.T) {
		svc := service.NewCalculatorService(&stubEvaluator{result: 7})
		result, err := svc.Calculate("2+5")
		require.NoError(t, err)
		require.Equal(t, 7.0, result)
	})

	t.Run("calculator service surfaces evaluator error", func(t *testing.T) {
		svc := service.NewCalculatorService(&stubEvaluator{err: errors.New("boom")})
		result, err := svc.Calculate("2+5")
		require.Error(t, err)
		require.Equal(t, 0.0, result)
	})

	t.Run("calculator service validates expression before evaluator", func(t *testing.T) {
		svc := service.NewCalculatorService(evaluator.NewEvaluator())
		result, err := svc.Calculate("2+")
		require.Error(t, err)
		require.Equal(t, 0.0, result)
	})
}
