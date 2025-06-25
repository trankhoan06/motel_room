package ginMail

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/modules/email/biz"
	"main.go/modules/email/model"
	"main.go/modules/email/storage"
	storageUser "main.go/modules/user/storage"
	"net/http"
)

func VerifyCodeEmail(db *gorm.DB, provider tokenprovider.TokenProvider) func(*gin.Context) {
	return func(c *gin.Context) {
		var verify model.VerifyAccountCode
		if err := c.ShouldBindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewLoginBiz(store, storeUser, provider, hash)
		token, err := business.NewVerifyEmail(c.Request.Context(), &verify, 60)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Status": true, "token": token})
	}
}
