package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/modules/user/biz"
	"main.go/modules/user/model"
	"main.go/modules/user/storage"
	"net/http"
)

func Login(db *gorm.DB, provider tokenprovider.TokenProvider) func(*gin.Context) {
	return func(c *gin.Context) {
		var l model.Login
		if err := c.ShouldBindJSON(&l); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewLoginBiz(store, provider, hash)
		token, err := business.NewLogin(c.Request.Context(), &l, 24*30)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
