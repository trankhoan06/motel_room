package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/config"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func CreateVerifyCodeEmail(db *gorm.DB, cfg *config.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		var verify model.VerifyAccountCode
		if err := c.ShouldBindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewSendEmailBiz(store, cfg)
		createVerify, err := business.NewCreateVerifyCodeEmail(c.Request.Context(), &verify, 60)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Status": true, "data": createVerify})
	}
}
