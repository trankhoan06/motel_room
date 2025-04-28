package ginRent

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/modules/rent/biz"
	"main.go/modules/rent/storage"
	"net/http"
)

func ListRent(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var user common.UserIdCommon
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewRentCommon(store)
		data, err := business.NewListRent(c.Request.Context(), user.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
