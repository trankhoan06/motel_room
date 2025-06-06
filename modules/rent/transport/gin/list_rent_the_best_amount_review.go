package ginRent

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/rent/storage"
	"net/http"
	"strconv"
)

func ListRentTheBestAmountReview(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		data, err := store.ListRentTheBestAmountReview(c.Request.Context(), limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
