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

func CreateVerifyCodeEmail(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var verify model.CreateVerifyAccount
		if err := c.ShouldBindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSqlModel(db)
		business := biz.NewSendEmailBiz(store, storeUser)
		createVerify, err := business.NewResendCodeEmail(c.Request.Context(), verify.Email, 5*60, verify.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Status": true, "data": createVerify})
	}
}
