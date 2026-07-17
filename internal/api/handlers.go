package api

import (
	"net/http"

	"calculator-backend/internal/models"
	"calculator-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func createCalculateHandler(calculatorService service.CalculatorService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.CalculateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, models.CalculateResponse{Error: "invalid request body"})
			return
		}

		result, err := calculatorService.Calculate(request.Expression)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CalculateResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.CalculateResponse{Result: result})
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
