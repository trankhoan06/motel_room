package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func ForgotPassword(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var forgot model.ForgotPassword
		if err := c.ShouldBindJSON(&forgot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewUserCommonBiz(store)
		verify, err := business.NewForgotPassword(c.Request.Context(), &forgot, 60)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": verify})
	}
}
