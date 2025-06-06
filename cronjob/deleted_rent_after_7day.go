package cronjob

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"main.go/modules/rent/storage"
	"net/http"
)

func DeletedRentAfter7day(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSqlModel(db)
		if err := store.DeletedAfter7Day(c.Request.Context()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Printf(err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
