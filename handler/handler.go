package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/jumia_task/database"
	"github.com/kjasuquo/jumia_task/services"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func (h *Handler) SearchNumberHandler(c *gin.Context) {
	country := c.Query("country")
	state := c.Query("state")

	customer, err := h.DB.ScanCustomerTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error from ScanCustomerTable"})
		return
	}

	numbers := services.GetAllValidatedNumbers(customer)

	num, err := h.DB.ScanNumberTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error from ScanNumberTable"})
		return
	}

	err = h.DB.InsertToNumbersTable(numbers, num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error Inserting number into Table"})
		return
	}

	search, err := h.DB.Search(country, state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error searching database"})
		return
	}

	c.JSON(http.StatusOK, search)

}
