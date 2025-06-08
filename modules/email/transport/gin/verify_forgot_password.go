package ginMail

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/modules/email/biz"
	"main.go/modules/email/model"
	"main.go/modules/email/storage"
	storageUser "main.go/modules/user/storage"
	"net/http"
)

func VerifyForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.VerifyAccountCode
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSqlModel(db)
		business := biz.NewSendEmailBiz(store, storeUser)
		if err := business.NewVerifyForgotPassword(c.Request.Context(), &data, 60*5); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
