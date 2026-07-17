package api

import (
	"calculator-backend/internal/middleware"
	"calculator-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(calculatorService service.CalculatorService) *gin.Engine {
	router := gin.New()
	router.Use(middleware.CORS())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Recovery())

	router.GET("/health", healthHandler)
	router.POST("/calculate", createCalculateHandler(calculatorService))

	return router
}
